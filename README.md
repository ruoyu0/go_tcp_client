# go_tcp_client

go语言 net包API：[http://studygolang.com/static/pkgdoc/pkg/net.htm](http://studygolang.com/static/pkgdoc/pkg/net.htm)

#### type Conn
    type Conn interface {
        // Read从连接中读取数据
        // Read方法可能会在超过某个固定时间限制后超时返回错误，该错误的Timeout()方法返回真
        Read(b []byte) (n int, err error)
        // Write从连接中写入数据
        // Write方法可能会在超过某个固定时间限制后超时返回错误，该错误的Timeout()方法返回真
        Write(b []byte) (n int, err error)
        // Close方法关闭该连接
        // 并会导致任何阻塞中的Read或Write方法不再阻塞并返回错误
        Close() error
        // 返回本地网络地址
        LocalAddr() Addr
        // 返回远端网络地址
        RemoteAddr() Addr
        // 设定该连接的读写deadline，等价于同时调用SetReadDeadline和SetWriteDeadline
        // deadline是一个绝对时间，超过该时间后I/O操作就会直接因超时失败返回而不会阻塞
        // deadline对之后的所有I/O操作都起效，而不仅仅是下一次的读或写操作
        // 参数t为零值表示不设置期限
        SetDeadline(t time.Time) error
        // 设定该连接的读操作deadline，参数t为零值表示不设置期限
        SetReadDeadline(t time.Time) error
        // 设定该连接的写操作deadline，参数t为零值表示不设置期限
        // 即使写入超时，返回值n也可能>0，说明成功写入了部分数据
        SetWriteDeadline(t time.Time) error
    }
* Conn接口代表通用的面向流的网络连接。多个线程可能会同时调用同一个Conn的方法。

#### func Dial
    func Dial(network, address string) (Conn, error)

* 在网络network上连接地址address，并返回一个Conn接口。可用的网络类型有：
* "tcp"、"tcp4"、"tcp6"、"udp"、"udp4"、"udp6"、"ip"、"ip4"、"ip6"、"unix"、"unixgram"、"unixpacket"
* 对TCP和UDP网络，地址格式是host:port或[host]:port，参见函数JoinHostPort和SplitHostPort。
    Dial("tcp", "12.34.56.78:80")
    Dial("tcp", "google.com:http")
    Dial("tcp", "[2001:db8::1]:http")
    Dial("tcp", "[fe80::1%lo0]:80")
* 对IP网络，network必须是"ip"、"ip4"、"ip6"后跟冒号和协议号或者协议名，地址必须是IP地址字面值。  
    Dial("ip4:1", "127.0.0.1")
    Dial("ip6:ospf", "::1")
* 对Unix网络，地址必须是文件系统路径。
