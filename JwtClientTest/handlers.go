package main

import (
	"github.com/gorilla/mux"
	"net/http"
	"fmt"
	"io/ioutil"
	"github.com/dgrijalva/jwt-go"
	"time"
)

var signingKey string = `|1|TYBNL92OcDy1ZziWR6BSK7Ybnjs=|9qlZ8sb7o3sqbF6oHzqeEsqJktk= ssh-rsa AAAAB3NzaC1yc2EAAAABIwAAAQEAq2A7hRGmdnm9tUDbO9IDSwBK6TbQa+PXYPCPy6rbTrTtw7PHkccKrpp0yVhp5HdEIcKr6pLlVDBfOLX9QUsyCOV0wzfjIJNlGEYsdlLJizHhbn2mUjvSAHQqZETYP81eFzLQNnPHt4EVVUh7VfDESU84KezmD5QlWpXLmvU31/yMf+Se8xhHTvKSCZIFImWwoG6mbUoWf9nzpIoaSjB+weqqUUmpaaasXVal72J+UX2B+2RPW3RcT0eOzQgqlJL3RKrTJvdsjE3JEAvGq3lGHSZXy28G3skua2SmVi/w4yCE6gbODqnTWlg7+wC604ydGXA8VJiS5ap43JXiUFFAaQ==
|1|IOoxp0K4z3xV9JylB7J/YxsQ/88=|V5iOMnLtFsqUUxPs+ngD84LP1ao= ssh-rsa AAAAB3NzaC1yc2EAAAABIwAAAQEAq2A7hRGmdnm9tUDbO9IDSwBK6TbQa+PXYPCPy6rbTrTtw7PHkccKrpp0yVhp5HdEIcKr6pLlVDBfOLX9QUsyCOV0wzfjIJNlGEYsdlLJizHhbn2mUjvSAHQqZETYP81eFzLQNnPHt4EVVUh7VfDESU84KezmD5QlWpXLmvU31/yMf+Se8xhHTvKSCZIFImWwoG6mbUoWf9nzpIoaSjB+weqqUUmpaaasXVal72J+UX2B+2RPW3RcT0eOzQgqlJL3RKrTJvdsjE3JEAvGq3lGHSZXy28G3skua2SmVi/w4yCE6gbODqnTWlg7+wC604ydGXA8VJiS5ap43JXiUFFAaQ==
|1|xtbQ6xCBDwT4o3O8GvIasNzyX8I=|1ZSP/vW+At2j0y38Ph9bmxcp3HM= ssh-rsa AAAAB3NzaC1yc2EAAAABIwAAAQEAq2A7hRGmdnm9tUDbO9IDSwBK6TbQa+PXYPCPy6rbTrTtw7PHkccKrpp0yVhp5HdEIcKr6pLlVDBfOLX9QUsyCOV0wzfjIJNlGEYsdlLJizHhbn2mUjvSAHQqZETYP81eFzLQNnPHt4EVVUh7VfDESU84KezmD5QlWpXLmvU31/yMf+Se8xhHTvKSCZIFImWwoG6mbUoWf9nzpIoaSjB+weqqUUmpaaasXVal72J+UX2B+2RPW3RcT0eOzQgqlJL3RKrTJvdsjE3JEAvGq3lGHSZXy28G3skua2SmVi/w4yCE6gbODqnTWlg7+wC604ydGXA8VJiS5ap43JXiUFFAaQ==
|1|HV6sxaxsEQfsP0dSH/IrSWhOVZ8=|7xrzWqz2fRKmcIHQwq6YV9gYgkA= ecdsa-sha2-nistp256 AAAAE2VjZHNhLXNoYTItbmlzdHAyNTYAAAAIbmlzdHAyNTYAAABBBMF5PEouqjURIo4iWblAVVivTXOEJOH2QRcCQvXFz2XHMrMyULPkzx2XAZ7EZy8n9xg79P3RH1qJgsgbncVenVg=
|1|3gSsWV7NtJWXbik3V2ynXguTgqA=|N/maNX3vPT/6dErjZe4Nx/sUEuY= ecdsa-sha2-nistp256 AAAAE2VjZHNhLXNoYTItbmlzdHAyNTYAAAAIbmlzdHAyNTYAAABBBGThjbzmqtFNr+AshTdg6hDhq6aWar9PdK0YVPh97HOq3e7bVYIWRUh1w415WglGX5ofclvgSkx1cewXVzUVD78=
|1|r69sSRr4PCK2ON0g7c4RwmKcb5E=|9Vy7ZtQz/yPzKkvB2jsT3c4LEHY= ecdsa-sha2-nistp256 AAAAE2VjZHNhLXNoYTItbmlzdHAyNTYAAAAIbmlzdHAyNTYAAABBBFcPHH1K5jRUWT1QN1b1c+yEAL90cYkEsKNzXy7O+TEm5COOO7K4I75ttJk2Wa7SyB6bvqMliUHwpkXcXcmi0e4=
|1|/6Fl1ZJwum3P8s9VVXtOKsAMbPM=|dBhtEEK5WZFHoKhqV/7e5wJQ6FQ= ecdsa-sha2-nistp256 AAAAE2VjZHNhLXNoYTItbmlzdHAyNTYAAAAIbmlzdHAyNTYAAABBBPPCdSE667Pln36Nz4qkFvoxlgJgS4tFdKRYVu+ZfvbFuE2qgJxADIk+8+UzaIx9e7fzeMdPJa/zkmRjW5aWRDg=
|1|KEqU9KyaXjFGAR70Kq6uk1LOHPk=|x9JMLZy25zwGwc8sfhXW34QpCUk= ecdsa-sha2-nistp256 AAAAE2VjZHNhLXNoYTItbmlzdHAyNTYAAAAIbmlzdHAyNTYAAABBBJZPvZ1aY1XJRJh28Zrjy9mo/MDLloKfnQEAJ26inzYtykvExhuymIQqMrpmUeTevskCn8tkihUTNbTtMWf4ELQ=
|1|iOyRz8f9/UFiTO2v6tdIiqYxpjc=|VIGPgptDDz1zgNvipoQppQYq2vk= ecdsa-sha2-nistp256 AAAAE2VjZHNhLXNoYTItbmlzdHAyNTYAAAAIbmlzdHAyNTYAAABBBE5Y0oLK1mnoQTRpmuDxWCKyIAfOk90P4sEHkud0gPblYJLwZPjK3u54gR78CcxYatg+SHehkQyQ11diw6u0VAM=
|1|ZXF63Ae/rb6xuSdsdDg440yYMPY=|68UOw28oyUoD2YKYY+SlxtVDgF0= ecdsa-sha2-nistp256 AAAAE2VjZHNhLXNoYTItbmlzdHAyNTYAAAAIbmlzdHAyNTYAAABBBO6GJUqonpl+EMcvD59hydInczFOOdTfs2rzl20fijJp242Km7xr8/qVH7FhwhT6KIA+vJhFWNszOkXan51aWFw=
|1|gJy/nT8U6/lkmV3h4fZk1BYQ7DA=|vFMgmuRd3h7c5gIl2DAPiiM5s3w= ecdsa-sha2-nistp256 AAAAE2VjZHNhLXNoYTItbmlzdHAyNTYAAAAIbmlzdHAyNTYAAABBBGDpFBLub54hGpVxfTfMjkQVZfbokBpLM/vuWp7jWcHcfIb4c9SiX1xsU8Apa/EAkf9/bk4w3dEB9gd44PAScmM=
|1|X591hycUfKfTnBITZch4pPSm03I=|OjnRbLkt0gg0OdxRJqaHAUbPUx8= ecdsa-sha2-nistp256 AAAAE2VjZHNhLXNoYTItbmlzdHAyNTYAAAAIbmlzdHAyNTYAAABBBHFHFZN5eRL/WSqHpwNnUEFektU6GgnsCEr9Y51rv0PUszyxUs/UQhhjYZdZ+WOyeRDditmwp/H86190ew7FgVE=
|1|QK+ahxEelNyNGRWUNNrYaZ1Of/c=|nfo1l7Q39MT+5v8AkZLtY0DnnGs= ecdsa-sha2-nistp256 AAAAE2VjZHNhLXNoYTItbmlzdHAyNTYAAAAIbmlzdHAyNTYAAABBBAgZdwHWVkIMtcHH2fjeiC1vDbe0hXuGB8/Rr/WrpJtrUsehp8O+YwiyYTaGCEAk2V8tRDZGDSvIkNsJyvnT1Uw=
|1|JTSUN30Tf4nrI42Ixo3Z9/BKMP4=|361Xhrv4GXA8wii8PdF5dbq9aq4= ecdsa-sha2-nistp256 AAAAE2VjZHNhLXNoYTItbmlzdHAyNTYAAAAIbmlzdHAyNTYAAABBBOE0GA3PJ/O88EnSUP+rJ0YClGzv2+IDvVGXAU+k67VjvaPUdQGPgufjr8tw4xfuHnXYW1X8C5kdj5fP8asIlVU=
|1|xE0T4/CQr5Ts1tWnKGSLpbJoPBg=|UR4zz06mgvKdp7qXMXzdWUxtpzY= ecdsa-sha2-nistp256 AAAAE2VjZHNhLXNoYTItbmlzdHAyNTYAAAAIbmlzdHAyNTYAAABBBCqX9waWHd8qPKtOHnuIn112e7w7gXF9F1Ec3OitX+ZD1bfhR6Yd02EW1ns8ylddWwkxfSGJGgUPZmvw8jsjxFs=
|1|CBDQGIjUWiqHFJaBiNjhLlCSsRg=|Pb8MLw447/lygirpUrUcVrxeP2w= ecdsa-sha2-nistp256 AAAAE2VjZHNhLXNoYTItbmlzdHAyNTYAAAAIbmlzdHAyNTYAAABBBJnCNgpOA5viPysEYcXULga+wksDNY30/MRYiyz/aXeCPbMUyjTocgMdAzYseM2oaK8/VeeBmvwFQlapbFPp3pE=
`

func addRoutes(r *mux.Router)  {
	r.HandleFunc("/test", UserRequest).Methods("GET")
}

func UserRequest(w http.ResponseWriter, r *http.Request)  {
	start := time.Now()

	r.ParseForm()
	email := r.FormValue("email")
	token, err := getToken(email)

	if err != nil {
		println(err)
		http.Error(w, http.StatusText(500), 500)
		return
	}

	response, err := http.Get(fmt.Sprintf("http://localhost:3002/test?email=%s&token=%s", email, token))
	if err != nil {
		println(err)
		http.Error(w, http.StatusText(500), 500)
		fmt.Println(time.Since(start))
		return
	}

	defer response.Body.Close()
	data, err := ioutil.ReadAll(response.Body)

	if err != nil {
		println(err)
		http.Error(w, http.StatusText(500), 500)
		fmt.Println(time.Since(start))
		return
	}

	tokenString := string(data)
	if tokenString != token {
		http.Error(w, http.StatusText(401), 401)
		fmt.Println(time.Since(start))
		return
	}

	w.Write([]byte("ok"))
	fmt.Println(time.Since(start))
}

func getToken(email string) (tokenString string, err error) {
    token := jwt.New(jwt.SigningMethodHS256)
	tokenString, err = token.SignedString([]byte(signingKey))

	return tokenString, err
}