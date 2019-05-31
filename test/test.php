<?php
$mc = new Memcached();
$mc->addServer('127.0.0.1', 7075);
$k = 'test';

$stime = explode(' ', microtime());
$num = 100000;

for($i=0;$i<$num;$i++){
    $id = $mc->get($k);
}

$etime = explode(' ',microtime());
$utime = $etime[0]+$etime[1]-($stime[0]+$stime[1]);

echo "use:{$utime} \n\n";
