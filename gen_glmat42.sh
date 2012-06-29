rm -rf glmat42
mkdir glmat42

for i in glmat33/*; do
        sed 's@chsc/gogl/gl33@chsc/gogl/gogl/gl42@g' "$i" > glmat42/`basename $i`
done

