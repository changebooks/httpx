# httpx
http + log
==

<pre>
profile, err := log.NewProfile(map[string]string{
    log.ProfileDirectory: "./log",
    log.ProfileChannel:   "http",
})

if err != nil {
    fmt.Println(err)
    return
}

stream, err := log.NewStream(profile)
if err != nil {
    fmt.Println(err)
    return
}

logger, err := log.NewLogger(stream, "test", 1)
if err != nil {
    fmt.Println(err)
    return
}

idRegister := &log.IdRegister{}
idRegister.SetTraceId("trace-id-10001")
idRegister.SetBizId("biz-id-20002")

urlBuilder := http.BaseUrlBuilder{}
baseUrl, err := urlBuilder.SetHost("https://www.baidu.com").Build()
if err != nil {
    fmt.Println(err)
    return
}

talkBuilder := http.TalkBuilder{}
talk, err := talkBuilder.SetBaseUrl(baseUrl).SetUseTrace(true).Build()
if err != nil {
    fmt.Println(err)
    return
}

http2, err := httpx.New(talk, logger)
if err != nil {
    fmt.Println(err)
    return
}

params := make(map[string]string)
params["wd"] = "abc"

req, url, err := http2.NewHttpRequest(idRegister, http.MethodGet, "s", params)
if err != nil {
    fmt.Println(url)
    fmt.Println(err)
    return
}

s := http2.Request(idRegister, req, 0)
fmt.Println(s.ToString())
if s.Elapsed != nil {
    fmt.Println(s.Elapsed.ToString())
}
</pre>
