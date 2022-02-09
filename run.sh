#go get github.com/pilu/fresh
sed -i -e "12 s/version: .*/version: $(date '+%Y.%m.%d.%H.%M')/" ./configs/development/app.yaml
fresh runner.conf