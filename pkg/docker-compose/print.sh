# 统一控制台输出
# shellcheck disable=SC2000,SC2046,SC2051,SC2059

rainbowColor=-2
rainbowColorIndex=0
rainbowColors=( 196 208 226 118 46 48 51 33 21 93 201 198 )
numRainbowColors=${#rainbowColors[@]}

foreground=38
background=48

# $3 The text
# $2 Color Code [0 - 255]  
# $1 Foreground 0 / Background 1
function print_colored_text {

	text=$1
	colorCode=$2

	if [[ $3 == 1 ]]
	then
		foregroundCode=$background
	else
		foregroundCode=$foreground
	fi

	if [[ $colorCode == -1 ]] 
	then
		printf "\033[0 m$text"
		return
	fi

	if [[ $colorCode -eq $rainbowColor ]]
	then
		for (( i=0 ; i < ${#text} ; i++ ))
		do 
			colorCode=${rainbowColors[$rainbowColorIndex]}
			rainbowColorIndex=$(($rainbowColorIndex +1))
			rainbowColorIndex=$(($rainbowColorIndex %$numRainbowColors))
			character=${text:i:1}
			printf "\033[${foregroundCode};5;${colorCode}m${character}\033[0m"
		done
		return
	fi

	printf "\033[${foregroundCode};5;${colorCode}m${text}\033[0m"
}

# Check if the $1 parameter is a number > 0 and <= $2
function is_in_range {
	if [[ "$1" =~ ^[0-9]+$ ]] && [ "$1" -ge 0 -a "$1" -le $2 ] && [ $# == 2 ]
	then
		echo $((1))
	else
		echo $((0))
	fi
}

# $1 = Red Component   [0 - 5]
# $2 = Green Component [0 - 5]
# $3 = Blue Component  [0 - 5]
function get_color_code {
	if [[ $(is_in_range $1 5) == 1 && $(is_in_range $2 5) == 1 && $(is_in_range $3 5) == 1 && $# == 3 ]]
	then
		echo $((16 + (36 * $1) + (6 * $2) + $3))
	else
		echo $((-1))
	fi
}

function print_usage {
	print_text "\n"
	print_text --rainbow "===================================="
	print_text -W        " print_text                                 "
	print_text --rainbow "   The most colorful echo command   "
	print_text --rainbow "===================================="

	print_text -Y "\n\nUsage:"
	print_text -W "print_text [[options] text]"
	print_text -W "Typical usage: " -G "print_text color1 text1 color2 text2 ..."

	print_text -Y "\n\nOptions:"
	print_text -C " -h|--help" 			-W " Print the correct usage" 
	print_text -C " -bk|--black" 		-W " Print the following parameters with a black color" -bk " example"
	print_text -C " -r|--red" 			-W " Print the following parameters with a red color" -r " example"
	print_text -C " -g|--green" 		-W " Print the following parameters with a green color" -g " example"
	print_text -C " -y|--yellow" 		-W " Print the following parameters with a yellow color" -y " example"
	print_text -C " -b|--blue" 			-W " Print the following parameters with a blue color" -b " example"
	print_text -C " -m|--magenta" 		-W " Print the following parameters with a magenta color" -m " example"
	print_text -C " -c|--cyan" 			-W " Print the following parameters with a cyan color" -c " example"
	print_text -C " -w|--white" 		-W " Print the following parameters with a white color" -w " example"
	print_text -C " -gy|--grey" 		-W " Print the following parameters with a grey color" -gy " example"
	print_text -C " -R|--Red"			-W " Print the following parameters with a bright red color" -R " example"
	print_text -C " -G|--Green"			-W " Print the following parameters with a bright green color" -G " example"
	print_text -C " -Y|--Yellow"		-W " Print the following parameters with a bright yellow color" -Y " example"
	print_text -C " -B|--Blue"			-W " Print the following parameters with a bright blue color" -B " example"
	print_text -C " -M|--Magenta"		-W " Print the following parameters with a bright magenta color" -M " example"
	print_text -C " -C|--Cyan"			-W " Print the following parameters with a bright cyan color" -C " example"
	print_text -C " -W|--White"			-W " Print the following parameters with a bright white color" -W " example"
	print_text -C " -Gy|--Gray"			-W " Print the following parameters with a bright grey color" -Gy " example"
	print_text -C " -h|--hex"			-W " Print the following parameters with a color from in a hex format.\n The hex color shall not start with the # but just 6 hex digits.\n example: " -G "print_text --hex F8E4B0 text \n"
	print_text -C " -rgb5"				-W " Print the following parameters with the specified R G B components.\n The R G B components must have a range between 0-5 and must be divided by a comma without spaces.\n example: " -G "print_text -rgb5 1,2,3 text \n"
	print_text -C " -rgb100"			-W " Print the following parameters with the specified R G B components.\n The R G B components must have a range between 0-100 and must be divided by a comma without spaces.\n example: " -G "print_text -rgb5 10,20,30 text \n"
	print_text -C " -rgb|-rgb255"		-W " Print the following parameters with the specified R G B components.\n The R G B components must have a range between 0-255 and must be divided by a comma without spaces.\n example: " -G "print_text -rgb5 100,200,255 text \n"
	print_text -C " -fg|--foreground"	-W " Color the foreground ( default )"
	print_text -C " -bg|--background"	-W " Color the background " -bg -R " example " -fg ""
	print_text -C " -n|-il|--in-line"	-W " With this option the command will not print the new line character after the execution\n"
	print_text -C " -ran|--random"		-W " Print the following parameters with a random color"
	print_text -C " -rb|--rainbow"		-W " Print the following parameters with a different color for each character following the classical Rainbow schema " -rb "example"
	print_text -C " -code|--color-code <code>" -W " Print the following parameters with the specified color code, the code must be a number between 0-255.\n"
	print_text -C " --get-color-code <r,g,b>" 	-W " Print the color code from the R G B components. The R G B components must have a range between 0-255 and must be divided by a comma without spaces.\n"



	print_text -Y "\n\nExamples:"
	print_text -il -G "print_text text" -W " will print: "
	print_text text
	
	print_text -il -G "print_text --Red RedText" -W " will print: "
	print_text --Red RedText
	
	print_text -il -G "print_text --Green GreenText" -W " will print: "
	print_text --Green GreenText
	
	print_text -il -G "print_text --Blue BlueText" -W " will print: "
	print_text --Blue BlueText
	
	print_text -il -G "print_text --Red \"RedText\" --Green \" GreenText\" --Blue \" BlueText\"" -W " will print: "
	print_text --Red RedText --Green " GreenText" --Blue " BlueText"

	print_text -il -G "print_text --rainbow \"Rainbow Colored Text\"" -W " will print: "
	print_text --rainbow "Rainbow Colored Text"
	
	print_text -il -G "print_text -rgb 100,200,30 \"Custom Color Text\"" -W " will print: "
	print_text -rgb 200,100,30 "Custom Color Text"

	print_text -il -G "print_text --hex FF88ee \"Color From Hex\"" -W " will print: "
	print_text --hex FF88ee "Color From Hex"
	
	print_text -il -G "print_text -B G -R o -Y o -B g -G l -R e" -W " will print: "
	print_text -B G -R o -Y o -B g -G l -R e

	print_text "\n"
	exit 1
}

function print_text() {    
  color=-1
  colorBackground=0
  inline=0
  
  while (( "$#" ))
  do
    case $1 in
      -h|--help)			print_usage ;;
        -bk|--black)		color=0 ;;
      -r|--red) 			color=1 ;;
      -g|--green) 		color=2 ;;
      -y|--yellow) 		color=3 ;;
      -b|--blue) 			color=4 ;;
      -m|--magenta) 		color=5 ;;
      -c|--cyan) 			color=6 ;;
      -w|--white) 		color=7 ;;
      -Gy|--Grey) 		color=8 ;;
      -R|--Red) 			color=9 ;;
      -G|--Green) 		color=10 ;;
      -Y|--Yellow) 		color=11 ;;
      -B|--Blue) 			color=12 ;;
      -M|--Magenta) 		color=13 ;;
      -C|--Cyan) 			color=14 ;;
      -W|--White) 		color=15 ;;
      -gy|--gray) 		color=16 ;;
      -fg|--foreground)	colorBackground=0 ;;
      -bg|--background) 	colorBackground=1 ;;
      -il|--in-line)		inline=1 ;;
      -ran|--random)		color=$(($RANDOM %255)) ;;
      -rb|--rainbow)		
        color=$rainbowColor
        rainbowColorIndex=0
      ;;
      -code|--color-code) 
        color=$2
        shift
      ;;
      --get-color-code) 
        IFS=, read r g b <<< "$2"
        if [[ $r && $g && $b ]]
        then
          r=$(($r/51))
          g=$(($g/51))
          b=$(($b/51))
          get_color_code $r $g $b
          exit 0
        else
          print_usage
        fi
      ;;
      -rgb5)
        IFS=, read r g b <<< "$2"
        if [[ $r && $g && $b ]]
        then
          color=$(get_color_code $r $g $b)
          shift
        else
          print_usage
        fi
      ;;
      -rgb100)
        IFS=, read r g b <<< "$2"
        if [[ $r && $g && $b ]]
        then
          r=$(($r/20))
          g=$(($g/20))
          b=$(($b/20))
          color=$(get_color_code $r $g $b)
          shift
        else
          print_usage
        fi
      ;;
      -rgb|-rgb255)
        IFS=, read r g b <<< "$2"
        if [[ $r && $g && $b ]]
        then
          r=$(($r/51))
          g=$(($g/51))
          b=$(($b/51))
          color=$(get_color_code $r $g $b)
          shift
        else
          print_usage
        fi
      ;;
      -H|--hex)
        if [[ $# -gt 1 && $2 =~ ^[0-9A-Fa-f]{6}$ ]]
        then
          hex=$2
          r=$(printf "%d\n" 0x${hex:0:2}) 
          g=$(printf "%d\n" 0x${hex:2:2})
          b=$(printf "%d\n" 0x${hex:4:2})
          r=$(($r/51))
          g=$(($g/51))
          b=$(($b/51))
          color=$(get_color_code $r $g $b)
          shift
        else
          print_usage
        fi
      ;;
      # print the text
        *) 
        text=$1
        print_colored_text "$text" $color $colorBackground 
      ;;
    esac
    shift
  done
  
  if [[ $inline == 0 ]]
  then
    printf "\n"
  fi
}

# 生成分割线
function split_line() {
    s=$1
    if [ "$1" == "" ]; then
        s="-"
    fi
    string=$s
    len=$(tput cols)
    for (( i = 1; i < $len; i++ )); do
        string=$string$s
    done
    echo "$string"
}

#---------------- 调用入口
function _title() { 
    print_text -bg -gy "$1"
    print_text -bg -gy $(split_line "-")
}

function _ok() {
    print_text -g "SUCCEED: $1"
}

# _ok 别名
function _succeed() {
    _ok "$@"
}

function _err() {
    print_text -r "FAILED: $1"
}

# _err 别名
function _failed() {
    _err "$@"
}

# _err 别名
function _error() {
    _err "$@"
}

function _warn() {
    print_text -y "WARNING: $1"
}

# _warn 别名
function _warning() {
    _warn "$@"
}

function _tip() {
    print_text -b "OUTPUT: $1"
}

# _tip 别名
function _info() {
    _tip "$@"
}
