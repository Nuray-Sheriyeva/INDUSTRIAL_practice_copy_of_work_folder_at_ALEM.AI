if [ -d "mystery" ]; then 
    BASE="mystery"
else
    BASE="."
fi

export INTERVIEW=$(head -n 179 "$BASE/streets/Buckingham_Place" | tail -n 1| cut -d "#" -f2)
echo "$INTERVIEW"
cat "$BASE/interviews/interview-$INTERVIEW"
echo "$MAIN_SUSPECT"