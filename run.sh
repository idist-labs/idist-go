#go get github.com/pilu/fresh
sed -i -e "13 s/version: .*/version: $(date '+%Y.%m.%d.%H.%M')/" ./configs/development/app.yaml
rm ./configs/development/app.yaml-e
fresh runner.conf