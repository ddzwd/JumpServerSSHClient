package terminal

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/signal"
	"syscall"
	"time"

	"JumpServerSSHClient/config"

	"JumpServerSSHClient/instance"
	"JumpServerSSHClient/server"

	"golang.org/x/crypto/ssh"
	"golang.org/x/crypto/ssh/terminal"
	"golang.org/x/term"
)

func GetAuthMethods(user *config.User) (auths []ssh.AuthMethod, err error) {
	if user.RsaKeyPath != "" {
		auths = append(auths, ssh.PublicKeysCallback(func() ([]ssh.Signer, error) {
			privateKeyBytes, err := ioutil.ReadFile("/Users/shanweijia/.ssh/id_rsa")
			if err != nil {
				instance.Logger.Fatalf("Failed to load private key: %v", err)
			}
			signer, err := ssh.ParsePrivateKey(privateKeyBytes)
			if err != nil {
				instance.Logger.Fatalf("Failed to parse private key: %v", err)
			}
			return []ssh.Signer{signer}, nil
		}))
	}
	auths = append(auths, ssh.PasswordCallback(func() (secret string, err error) {
		if user.Password == "" {
			// 当密码为空时,需要提示用户输入密码
			fmt.Printf("请输入密码:")
			fmt.Scanln(&secret)
		} else {
			secret = user.Password
		}
		return
	}))
	auths = append(auths, ssh.KeyboardInteractiveChallenge(func(name, instruction string, questions []string, echos []bool) (answers []string, err error) {
		for _, question := range questions {
			if question == "[MFA auth]: " {
				if user.SecretKey != "" {
					answers = append(answers, server.NewMFA(user.SecretKey))
				} else {
					instance.Logger.Warn("user %s need MFA auth, but SecretKey is not set. add this to auto complete", user.UniqueId)
				}
			} else {
				// 提示question
				fmt.Printf("%s ", question)
				// 获取输入
				var answer string
				fmt.Scanln(&answer)
				answers = append(answers, answer)
			}
		}
		return
	}))
	return
}

// 启动ssh终端
func StartTerminal(user *config.User) {
	auths, err := GetAuthMethods(user)
	if err != nil {
		instance.Logger.Fatalf("Failed to get auth methods: %v", err)
	}
	// 创建ssh客户端
	sshConfig := &ssh.ClientConfig{
		User:            user.UserName,
		Auth:            auths,
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
		Timeout:         10 * time.Second,
	}
	sshConfig.SetDefaults()
	sshConfig.Ciphers = []string{"chacha20-poly1305@openssh.com"}
	sshConfig.KeyExchanges = []string{"curve25519-sha256@libssh.org"}
	sshConfig.MACs = []string{"hmac-sha2-256-etm@openssh.com"}

	conn, err := ssh.Dial("tcp", fmt.Sprintf("%s:%s", user.Host, user.Port), sshConfig)
	if err != nil {
		instance.Logger.Fatalf("Failed to dial: %s\n", err)

	}
	defer conn.Close()
	session, err := conn.NewSession()
	if err != nil {
		instance.Logger.Fatalf("Failed to create session: %s\n", err)
	}
	defer session.Close()

	// 启用伪终端
	fd := int(os.Stdin.Fd())
	state, err := terminal.MakeRaw(fd)
	if err != nil {
		instance.Logger.Fatal(err)
	}
	defer terminal.Restore(fd, state)

	width, height, err := term.GetSize(fd)
	if err != nil {
		instance.Logger.Fatal(err)
	}

	modes := ssh.TerminalModes{
		ssh.ECHO:          1,
		ssh.TTY_OP_ISPEED: 14400,
		ssh.TTY_OP_OSPEED: 14400,
	}

	if err := session.RequestPty("xterm-256color", height, width, modes); err != nil {
		session.Close()
		panic("request for pseudo terminal failed: " + err.Error())
	}

	// 处理终端大小调整信号
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGWINCH)
	defer signal.Stop(sigChan)
	done := make(chan struct{})
	go func() {
		for {
			select {
			case <-sigChan:
				width, height, err := term.GetSize(fd)
				if err == nil {
					session.WindowChange(height, width)
				}
			case <-done:
				return
			}
		}
	}()

	session.Stdin = os.Stdin
	session.Stdout = os.Stdout
	session.Stderr = os.Stderr

	err = session.Shell()
	if err != nil {
		instance.Logger.Fatalf("Failed to start shell: %s\n", err)
	}

	if err := session.Wait(); err != nil {
		instance.Logger.Fatalf("session colsed: %s\n", err)
	}
	done <- struct{}{}

}
