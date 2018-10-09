package thirdparty

func GetSeelogConfigContent() string {
	return `
<seelog type="sync" minlevel="trace" maxlevel="critical">
    <outputs formatid="fullmsg">
        <filter levels="trace,debug,info,warn">
            <rollingfile type="size" filename="./debug.log" maxrolls="5" maxsize="1048576"/>
        </filter>
        <filter levels="error,critical">
            <rollingfile type="size" filename="./error.log" maxrolls="5" maxsize="1048576"/>
        </filter>
    </outputs>
    <formats>
        <format id="fullmsg" format="%Date %Time [%LEV] [%File:%Line] [%Func] %Msg%n"/>
    </formats>
</seelog>
`
}