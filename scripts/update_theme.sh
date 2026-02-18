if [ $# -eq 0 ]; then
	echo "Usage: update_theme.sh <path to picture>"
	exit 1
fi

pipx run pywal -i "$1"
