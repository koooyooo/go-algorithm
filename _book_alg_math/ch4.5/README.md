## Q 4.5.1

| # | A | B | C | D |
| --- | --- | --- | --- | --- |
| 分類 | 重みなし無向グラフ | 重み付き無向グラフ | 重みなし有向グラフ | 重み付き有向グラフ |
 | 次数最大の頂点 | 1 | 5 | 3 | 2 |

## Q 4.5.2

###  E: 
- 経路：存在しない  
### F: 
- 経路：存在する（オイラーグラフ)  
- 例示： 1-2-3-4-5-9-8-7-6-10-11-12-13-14-9-13-8-12-7-11-6-9-4-8-3-7-2-6-1

## Q 4.5.3
### 1:
グラフを最小単位で見ると、菱形の中央に縦の対角線が入っている
- 上部の頂点と下部の頂点に注目した場合、
  - 対角線を避けると 2Stepで下部の色は上部と同色になる必要がある
  - 対角線を通るなら 1Stepで下部の色は上部と異色になる必要がある
この2点に矛盾を抱えているため、2色で塗り分ける経路は存在しない

### 2
横軸のラインが同色になるように、赤・青・緑を横縞で構成すれば3色で塗り分け可能

## 4.5.4
1, 5, 2, 10, 8, 13, 12, 6, 14, 7, 3, 11, 4, 9

## 4.5.5