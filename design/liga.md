# Liga 西甲联赛

## 赛程

预测新赛季赛程

    1 获取新赛季球队名单，未定的用A/B等字母替代
        地址：http://www.laliga.es/en/laliga-santander
        标签和属性：
            <div class="container" id="p15042018050101">
                <div id="equipos">
                    <div class="columna laliga-santander">
                        <span class="nombre"> -- text
    2 制定对赛规则
        每一轮：一次抽一队，前面十次随机抽出的10队作为主队，然后剩下10队中随机抽出的一队依次作为前面10队的对手；已经存在的对阵，下次抽取时，不可再出现；每个队都只有19个主场和19个客场，每两个队都只能相遇两次，分别是在主场一次；
        动态规划：上一轮的结果会影响到下一轮的抽签，比如，上一轮已经抽出过的对阵，以后就不能再出现了；也就是主队A已经抽到过客队B，接下来就不能再抽到了；换句话说，在A出现在主队的位置后，B队就不在待抽队列里了；
    3 制定赛程
