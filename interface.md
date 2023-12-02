# 接口文件
### 此文件敘述各個 Package 的開放接口 (public function) 和內容說明

--- 
## Packages 索引

###  - [failover](#failover)

###  - [index](#index)

###  - [kvstore](#kvstore)

###  - [kvstorekey](#kvstorekey)

###  - [proxy](#proxy)

###  - [route](#route)

###  - [secretutil](#secretutil)

###  - [setting](#setting)

###  - [teapot](#teapot)

###  - [tools](#tools)

###  - [teapot_maker](#teapot_maker)

###  - [template](#template)

###  - [tunnel](#tunnel)

###  - [util](#util)

--- 
## failover
``` go
func New(store kvstore.Store) *Failover 
```
* 這個 Func 沒有註解

``` go
func (f *Failover) UpdateTunnelIpAsPossible(attempts uint, delay time.Duration) 
```
* 這個 Func 沒有註解

``` go
func (f *Failover) UpdateAll() 
```
* 這個 Func 沒有註解

--- 
## index
``` go
func New(dbPath string, store kvstore.Store, client *http.Client) *Index 
```
* 這個 Func 沒有註解

``` go
func (index *Index) GetIndexVersion() int 
```
* 這個 Func 沒有註解

``` go
func (index *Index) FetchIndicesIfNeed() 
```
* 這個 Func 沒有註解

``` go
func (index *Index) FetchIndices() 
```
* 這個 Func 沒有註解

``` go
func (index *Index) IsContainIndex(domainHost string) bool 
```
* 這個 Func 沒有註解

--- 
## kvstore
``` go
func NewKvStore(path string, platform Platform) (*KvStore, error) 
```
* 這個 Func 沒有註解

``` go
func (kvStore *KvStore) GetRandomFromList(key kvstorekey.List) (string, error) 
```
* 這個 Func 沒有註解

``` go
func (kvStore *KvStore) GetStringList(key kvstorekey.List) ([]string, error) 
```
* 這個 Func 沒有註解

``` go
func (kvStore *KvStore) PutStringList(key kvstorekey.List, value []string) error 
```
* 這個 Func 沒有註解

``` go
func (kvStore *KvStore) GetString(key kvstorekey.Single) (string, error) 
```
* 這個 Func 沒有註解

``` go
func (kvStore *KvStore) PutString(key kvstorekey.Single, value string) error 
```
* 這個 Func 沒有註解

``` go
func (kvStore *KvStore) DeleteString(key kvstorekey.Single) error 
```
* 這個 Func 沒有註解

``` go
func (kvStore *KvStore) DeleteStringList(key kvstorekey.List) error 
```
* 這個 Func 沒有註解

--- 
## proxy
``` go
func New(addr string, routeLookup func(domain string) (string, error)) *httpProxy 
```
* 這個 Func 沒有註解

``` go
func (p *httpProxy) SetTunnel(tunnel *tunnel.Tunnel) 
```
* 這個 Func 沒有註解

``` go
func (p *httpProxy) SetDialerForTest(dialer Dialer) 
```
* 這個 Func 沒有註解

--- 
## route
``` go
func (r *Route) RouteLookup(domain string) (string, error) 
```
* 這個 Func 沒有註解

``` go
func (r *Route) GetRouteCacheCount(domain string) int 
```
* 這個 Func 沒有註解

``` go
func (r *Route) RefreshCache() 
```
* 這個 Func 沒有註解

``` go
func (r *Route) DeleteAllCache() bool 
```
* 這個 Func 沒有註解

--- 
## secretutil
``` go
func GenerateBBForwarded(ip string) string 
```
* 這個 Func 沒有註解

``` go
func EncodeJsonArr(strs []string) string 
```
* 這個 Func 沒有註解

``` go
func DecodeJsonArr(str string) ([]string, error) 
```
* 這個 Func 沒有註解

``` go
func EncodeRc4(plainText, key string) (string, error) 
```
* 這個 Func 沒有註解

``` go
func DecodeRc4(hexCipherText, key string) (string, error) 
```
* 這個 Func 沒有註解

``` go
func DecodeTeapotContent(s string) (string, error) 
```
* 這個 Func 沒有註解

``` go
func ExtractTeapotCipher(input string) string 
```
* 這個 Func 沒有註解

--- 
## setting
``` go
func UpdateCountry(fetcher *util.Fetcher, store kvstore.Store) error 
```
* 這個 Func 沒有註解

``` go
func (s *Setting) GetCountry() string 
```
* 這個 Func 沒有註解

``` go
func NewSetting(store kvstore.Store, client *http.Client, lang SysLang) *Setting 
```
* 這個 Func 沒有註解

``` go
func (s *Setting) UpdateCountryAsPossible(attempts uint, delay time.Duration, coolDown time.Duration) 
```
* 這個 Func 沒有註解

``` go
func (s *Setting) UpdateAll() 
```
* 這個 Func 沒有註解

--- 
## teapot
``` go
func NewTeapot(kvStore kvstore.Store, client *http.Client) (*Teapot, error) 
```
* 這個 Func 沒有註解

``` go
func (teapot *Teapot) SwitchToFirstAvailableBackupDomain() error 
```
* 這個 Func 沒有註解

--- 
## tools
``` go
func GetTeapotXmlString(hosts []string) (string, error) 
```
* 這個 Func 沒有註解

``` go
func MakeTeapotXmlFile(hosts []string, path string) error 
```
* 這個 Func 沒有註解

``` go
func MakeTeapotHtmlFile(hosts []string, templatePath string, outputPath string) error 
```
* 這個 Func 沒有註解

``` go
func GetKey() []byte 
```
* 這個 Func 沒有註解

--- 
## teapot_maker
``` go
func GetTeapotXmlString(hosts []string) (string, error) 
```
* 這個 Func 沒有註解

``` go
func MakeTeapotXmlFile(hosts []string, path string) error 
```
* 這個 Func 沒有註解

``` go
func MakeTeapotHtmlFile(hosts []string, templatePath string, outputPath string) error 
```
* 這個 Func 沒有註解

``` go
func GetKey() []byte 
```
* 這個 Func 沒有註解

--- 
## tunnel
``` go
func NewManager(t *Tunnel) *Manager 
```
* 這個 Func 沒有註解

``` go
func (m *Manager) StartTunnelAsPossible(coolDown time.Duration) 
```
* 這個 Func 沒有註解

``` go
func New(setting *Setting, db kvstore.Store, onUpdate func(status int)) *Tunnel 
```
* 這個 Func 沒有註解

``` go
func (tunnel *Tunnel) Close() 
```
* 這個 Func 沒有註解

``` go
func (tunnel *Tunnel) GetSshDialer() Dialer 
```
* 這個 Func 沒有註解

``` go
func (tunnel *Tunnel) IsConnected() bool 
```
* 這個 Func 沒有註解

``` go
func (tunnel *Tunnel) IsDisabled() bool 
```
* 這個 Func 沒有註解

``` go
func (tunnel *Tunnel) GetSecondSinceConnected() float64 
```
* 這個 Func 沒有註解

``` go
func (tunnel *Tunnel) GetBBForwarded() string 
```
* 這個 Func 沒有註解

``` go
func (tunnel *Tunnel) GetStartUpDuration() string 
```
* 這個 Func 沒有註解

--- 
## util
``` go
func NewFetcher(client *http.Client) *Fetcher 
```
* 這個 Func 沒有註解

``` go
func (f *Fetcher) Fetch(u *url.URL) (string, error) 
```
* 這個 Func 沒有註解

