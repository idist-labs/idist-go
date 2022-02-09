read_var() {
  if [ -z "$1" ]; then
    echo "environment variable name is required"
    return 1
  fi

  local ENV_FILE='.env'
  if [ ! -z "$2" ]; then
    ENV_FILE="$2"
  fi

  local VAR=$(grep $1 "$ENV_FILE" | xargs)
  IFS="=" read -ra VAR <<< "$VAR"
  echo ${VAR[1]}
}

endpoint=$(read_var REGISTRY)
sed -i -e "13 s/version: .*/version: $(date '+%Y.%m.%d.%H.%M')/" ./configs/production/app.yaml
rm ./configs/production/app.yaml-e
sed -i -e "11 s/version: .*/version: $(date '+%Y.%m.%d.%H.%M')/" ./configs/development/app.yaml
rm ./configs/development/app.yaml-e

echo "Building and pushing dockerfile to $endpoint..."
docker build -t $endpoint . && docker push $endpoint