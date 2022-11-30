set title "State Trie - Path Types"
set key off
set datafile separator ","; 
set ylabel "%" offset 1
set yrange [0:55]; 
set style fill solid; 
set boxwidth 0.5; 
set xtics nomirror rotate
set tic scale 0
set grid ytics
unset border
plot "statetrie_pathtypes.csv" using 0:2:xtic(1) with boxes