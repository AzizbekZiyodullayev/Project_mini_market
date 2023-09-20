#!/bin/bash

# -------- Task 1 --------

# EASY
# 'Grep' -A:
  # grep -Eiw 'the|that|then|those'
# 'Grep' -B:
  # grep "\([0-9]\) *\1"
# 'Sed' command #3
  # sed  's/\(thy\)/\{\1\}/gi'

# MEDIUM
#'Sed' command#1
    #sed  's/ the / this /1'
#'Sed' command#2
    # sed  's/\(thy\)/your/gi'
#'Grep' command#1
    # grep ' the ' || grep -w 'the'
#'Grep' command#2
    # grep -wi 'the'
#'Grep' command#3
    # grep -viw 'that'
#'Awk' command#1
    # awk '{if ($4<=0) print "Not all scores are available for",$1}'
#'Awk' command#2
    # awk '{if ($2<50 || $3<50 || $4<50){ print $1,": Fail"} else {print $1,": Pass"}}'
#'Awk' command#3
    # awk '{
    #   if (($2+$3+$4)/3>80){
    #     print $1,$2,$3,$4" : A"
    #   } else if (($2+$3+$4)/3>60) {
    #      print $1,$2,$3,$4" : B"
    #   } else if (($2+$3+$4)/3>50) {
    #      print $1,$2,$3,$4" : C"
    #   }else{
    #     print $1,$2,$3,$4" : FAIL"
    #   }}'
#'Awk' command#4
# awk 'BEGIN {
#         RS=""
#         FS="\n"
#         OFS=""
#     }
#     {
#       a=1
#       while (a<NF)
#           {print $a,";"$(a+1)
#         a=a+2}
#     }'

# --------- Task 2 --------

#'Cut' #1
# cut -b 3 || cut -c 3

#'Cut' #2
# cut -b 2,7 || cut -c 2,7

#'Cut' #3
# cut -b 2-7 || cut -c 2-7

#'Cut' #4
# cut -b -4 || cut -c -4

#'Cut' #5
# cut -f -3

#'Cut' #6
# cut -c 13-

#'Cut' #7
# cut -d ' ' -f 4

#'Cut' #8
# cut -d ' ' -f -3

#'Cut' #9
# cut -f 2-

#'Head' #1
# head -n 20

#'Head' #2
# head -c 20

#'Tail' #1
# tail -n 20

#'Tail' #1
# tail -c 20

#Middle of a text
# head -n 22 | tail -n 11

#'Tr' #1
# tr '()' '[]'

#'Tr' #2
# tr -d [a-z]

#'Tr' #3
# tr -s ' '

#'Sort' #1
# sort

#'Sort' #2
# sort -r

#'Sort' #3
# sort -n

#'Sort' #4
# sort -nr

#'Sort' #5
# sort -t $'\t' -k 2 -nr

#'Sort' #6
# sort -t $'\t' -k 2 -n

#'Sort' #6
# sort -t $'|' -k 2 -nr

#'Uniq' #1
# uniq

#'Uniq' #2
# uniq -c | cut -c 7-

#'Uniq' #3
# uniq -ci | cut -c 7-

#'Uniq' #4
# uniq -u

#'Read in an Array'
# while read line; do
#     countries=(${countries[@]} $line)
# done
# echo ${countries[@]}

# 'Slice an Array'
# while read line; do
#     countries=(${countries[@]} $line)
# done
# echo ${countries[@]:3:5}

# 'Concatenate an array with itself'
# while read line; 
# do
#     countries=(${countries[@]} $line)
# done
# countries=(${countries[@]} ${countries[@]} ${countries[@]})

# echo ${countries[*]}

# 'Display an element of an array'
# while read line; 
# do
#     countries=(${countries[@]} $line)
# done
# echo ${countries[3]}

# 'Count the number of elements in an Array'
# while read line; 
# do
#     countries=(${countries[@]} $line)
# done
# echo ${#countries[@]}

# 'Filter an Array with Patterns'
# while read line; 
# do
#     if [[ $line != *[aA]* ]]; then
#         echo $line
#     fi

# done

# 'Remove the First Capital Letter from Each Element'
# while read line; 
# do
#     line=.${line#[A-Z]}
#     countries=(${countries[@]} $line)
# done
# echo ${countries[@]}