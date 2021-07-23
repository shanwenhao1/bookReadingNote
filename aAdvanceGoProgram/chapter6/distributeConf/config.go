package distributeConf

//
//import (
//	"encoding/json"
//	"golang.org/x/net/context"
//	"time"
//)
//
//var configPath = `/configs/remote_config.json`
//var kApi client.KeysAPI
//
//type ConfigStruct struct {
//	Addr           string `json:"addr"`
//	AesKey         string `json:"aes_key"`
//	HTTPS          bool   `json:"https"`
//	Secret         string `json:"secret"`
//	PrivateKeyPath string `json:"private_key_path"`
//	CertFilePath   string `json:"cert_file_path"`
//}
//
//var appConfig ConfigStruct
//
//func init() {
//	cfg := client.Config{
//		Endpoints:               []string{"http://127.0.0.1:2379"},
//		Transport:               client.DefaultTransport,
//		HeaderTimeoutPerRequest: time.Second,
//	}
//
//	c, err := client.New(cfg)
//	if err != nil {
//		log.Fatal(err)
//	}
//	kApi = client.NewKeysAPI(c)
//	initConfig()
//}
//
//func watchAndUpdate() {
//	w := kApi.Watcher(configPath, nil)
//	go func() {
//		// watch 该节点下的每次变化
//		for {
//			resp, err := w.Next(context.Background())
//			if err != nil {
//				log.Fatal(err)
//			}
//			log.Println("new values is ", resp.Node.Value)
//
//			err = json.Unmarshal([]byte(resp.Node.Value), &appConfig)
//			if err != nil {
//				log.Fatal(err)
//			}
//		}
//	}()
//}
//
//func initConfig() {
//	resp, err = kApi.Get(context.Background(), configPath, nil)
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	err := json.Unmarshal(resp.Node.Value, &appConfig)
//	if err != nil {
//		log.Fatal(err)
//	}
//}
//
//func getConfig() ConfigStruct {
//	return appConfig
//}
//
//func main() {
//	// init your app
//}
