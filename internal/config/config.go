package config

import(
	"os"
	"encoding/json"
)
const(
	configfilename = "/.gatorconfig.json"
)
type Config struct {
	DB_Url				string
	CurrentUserName 	string
}

func Read() (Config, error){

	filepath, err := os.UserHomeDir()
	if err != nil {
		return Config{}, err
	}
	fullpath := filepath+configfilename

	data, err := os.ReadFile(fullpath)
	if err != nil {
		return Config{}, err
	}

	var cfg Config
	err = json.Unmarshal(data,&cfg)
	if err != nil {
		return Config{}, err
	}
	return cfg,nil
}

func (c *Config)SetUser(user string) error{
	c.CurrentUserName = user

	filepath, err := os.UserHomeDir()
	if err != nil {
		return err
	}

	fullpath := filepath+configfilename
	file, err := os.Create(fullpath)
	if err != nil {
		return err
	}

	encoder := json.NewEncoder(file)
	err = encoder.Encode(c)
	if err != nil {
		return err
	}
	return nil
}



