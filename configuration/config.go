package configuration



import ("fmt" 
"github.com/spf13/viper")


type Config struct{
	Prod string 
}
var AppConfig *Config

func LoadConfig(){

	var cfg Config

	viper.SetConfigFile("config.yml")

	viper.AddConfigPath(".")


	if err:=viper.ReadInConfig(); err!=nil{ 
		panic(fmt.Errorf("The following error occured: %w",err))
	}

    if err:=viper.Unmarshal(&cfg);err!=nil{
		panic(fmt.Errorf("The following error occured: %w",err))
		

	}    


	AppConfig=&cfg
	
	









}