package tomlConfig

import(
    "os"
    "fmt"
    "path/filepath"
    "github.com/BurntSushi/toml"
)

func Load(paths[]string, conf interface{})(error){
    // Directory where the binary is located. This is useful if we are testing.
    rootDir, _ := filepath.Abs(filepath.Dir(os.Args[0]))
    confFiles := append([]string{
        filepath.Join(rootDir, "settings.mine.toml"),
        // This is mainly used as an example file. Most of the time, mine.toml
        // should be loaded
        filepath.Join(rootDir, "settings.toml"),
    }, paths...)

    // Attempt to load all of the configs above.
    for _,f := range confFiles{
        _, err := toml.DecodeFile(f, conf)
        if(err!=nil){
            continue
        }
        return err
    }

    // If we reach this point then
    err := fmt.Errorf("Unable to load configuration file. Attempted dirctory %v", confFiles)
    return err;
}

func Create(filename string, conf interface{})error{
    file, err := os.Create(filename)
    if err != nil{
        return err
    }
    defer file.Close()

    encoder := toml.NewEncoder(file)
    encoder.Encode(conf)

    return nil
}
