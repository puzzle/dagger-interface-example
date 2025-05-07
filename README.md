Example code from https://docs.dagger.io/api/interfaces/

Calling `dagger functions` results in:

```
./dagger.gen.go:161:21: too many return values                
        have (bool, nil)                                      
        want (bool)                                           
./dagger.gen.go:166:19: too many return values                
        have (bool, error)                                    
        want (bool)                                           
./dagger.gen.go:166:29: undefined: ctx    
