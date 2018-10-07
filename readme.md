# Basic Static File Server
an incredibly simple, almost hello-worldish server in go.  Simple serves the current working directory at 
`localhost:3000` unless the path or port arguments are parsed when the server is started. 

## command line args
- `-port=3000` set the port
- `-path=public` set the path relative to the current dir.

NB: no license, use at own risk. 
NB2: you'll need to compile this and place somewhere (ie: your bin dir) where its available in your path.