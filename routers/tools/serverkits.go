package tools

import (
	"Plug-Ins/databases/mysql"
	"fmt"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/ssh"
	"net"
	"strconv"
	"time"
)

// ip :"81.68.254.93"
// username : "ubuntu"
// password :"fly0203*"
// port: 22

func GetServeList(ctx *gin.Context) {
	res := mysql.SelectAllMysql("select * from  kits_server")
	for _, value := range res {
		delete(value, "server_pwd")
		delete(value, "server_account")
		delete(value, "server_port")
		value["status"] = "--"
		value["cpu_occupancy"] = "--"
		value["memory"] = "--"
		value["tags"] = "--"
		value["systeminfo"] = "--"
	}
	ctx.JSON(200, gin.H{
		"data": res,
	})
}

func GetServerInfo(ip string, username string, password string, port string) ServerInfo {
	s1, _ := strconv.Atoi(port)
	cli := New(ip, username, password, s1)
	cpuinfo, err := cli.Run("cat /proc/cpuinfo")
	meminfo, err := cli.Run("cat /proc/meminfo")
	diskuse, err := cli.Run("df -h")
	if err != nil {
		panic("1")
	}
	//err = cli.client.Close()
	//if err != nil {
	//	return ServerInfo{}
	//}
	return ServerInfo{
		CpuInfo:  GetCpuInfo(cpuinfo),
		MemInfo:  GetMemInfo(meminfo),
		DataList: GetDataList(diskuse),
	}

}

// Cli GetServeStatus SSH获取服务器实时状态
type Cli struct {
	IP         string      //IP地址
	Username   string      //用户名
	Password   string      //密码
	Port       int         //端口号
	client     *ssh.Client //ssh客户端
	LastResult string      //最近一次Run的结果
}

// New 创建命令行对象
// @param ip IP地址
// @param username 用户名
// @param password 密码
// @param port 端口号,默认22
func New(ip string, username string, password string, port ...int) *Cli {
	cli := new(Cli)
	cli.IP = ip
	cli.Username = username
	cli.Password = password
	if len(port) <= 0 {
		cli.Port = 22
	} else {
		cli.Port = port[0]
	}
	return cli
}

// Run 执行shell
// @param shell shell脚本命令
func (c Cli) Run(shell string) (string, error) {
	if c.client == nil {
		if err := c.connect(); err != nil {
			return "", err
		}
	}
	session, err := c.client.NewSession()
	if err != nil {
		return "", err
	}
	defer session.Close()
	buf, err := session.CombinedOutput(shell)

	c.LastResult = string(buf)
	return c.LastResult, err
}

// GetServerInfo 返回单条服务器数据

// 连接

func (c *Cli) connect() error {
	config := ssh.ClientConfig{
		User: c.Username,
		Auth: []ssh.AuthMethod{ssh.Password(c.Password)},
		HostKeyCallback: func(hostname string, remote net.Addr, key ssh.PublicKey) error {
			return nil
		},
		Timeout: 10 * time.Second,
	}
	addr := fmt.Sprintf("%s:%d", c.IP, c.Port)
	sshClient, err := ssh.Dial("tcp", addr, &config)
	if err != nil {
		return err
	}
	c.client = sshClient
	return nil
}

//func (c *Cli) RunTerminal(shell string, stdout, stderr io.Writer) error {
//	if c.client == nil {
//		if err := c.connect(); err != nil {
//			return err
//		}
//	}
//	session, err := c.client.NewSession()
//	if err != nil {
//		return err
//	}
//	defer session.Close()
//
//	fd := int(os.Stdin.Fd())
//	oldState, err := terminal.MakeRaw(fd)
//	if err != nil {
//		panic(err)
//	}
//	defer terminal.Restore(fd, oldState)
//
//	session.Stdout = stdout
//	session.Stderr = stderr
//	session.Stdin = os.Stdin
//
//	termWidth, termHeight, err := terminal.GetSize(fd)
//	if err != nil {
//		panic(err)
//	}
//	// Set up terminal modes
//	modes := ssh.TerminalModes{
//		ssh.ECHO:          1,     // enable echoing
//		ssh.TTY_OP_ISPEED: 14400, // input speed = 14.4kbaud
//		ssh.TTY_OP_OSPEED: 14400, // output speed = 14.4kbaud
//	}
//
//	// Request pseudo terminal
//	if err := session.RequestPty("xterm-256color", termHeight, termWidth, modes); err != nil {
//		return err
//	}
//
//	session.Run(shell)
//	return nil
//}
