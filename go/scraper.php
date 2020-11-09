<?php
echo "hello world!!";

$url = "https://show.1688.com/pinlei/industry/pllist.html?spm=a260j.12536015.jr60bfo3.1.10f3700ecaYFQt&&sceneSetId=857&sceneId=2158&bizId=4404";
$ch = curl_init();
curl_setopt($ch, CURLOPT_URL, $url);
curl_setopt($ch, CURLOPT_HEADER, TRUE);
curl_setopt($ch, CURLOPT_RETURNTRANSFER, TRUE);
$data = curl_exec($ch);
$httpCode = curl_getinfo($ch, CURLINFO_HTTP_CODE);
curl_close($ch);

print_r($data);