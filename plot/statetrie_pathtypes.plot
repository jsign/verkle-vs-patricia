set title "State Trie - Path Types"
set key off
set datafile separator ","; 
set ylabel "%" offset 1
set style fill solid; 
set boxwidth 0.5; 
set xtics rotate
plot "statetrie_pathtypes.csv" using 0:2:xtic(1) with boxes