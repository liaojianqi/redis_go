// implemention map(ordered) use skiplist
#include <cstdio>
#include <string>
#include <iostream>
using namespace std;

#define ZSKIPLIST_MAXLEVEL 32 /* Should be enough for 2^32 elements */
#define ZSKIPLIST_P 0.25      /* Skiplist P = 1/4 */
int zslRandomLevel(void) {
    int level = 1;
    while ((random()&0xFFFF) < (ZSKIPLIST_P * 0xFFFF))
        level += 1;
    return (level<ZSKIPLIST_MAXLEVEL) ? level : ZSKIPLIST_MAXLEVEL;
}

struct SkipNode {
    int score, value, level;
    struct Level {
        int span;
        SkipNode *next;
    } levels[];
};
struct SkipList {
    SkipNode *head, *tail;
    int length, level;
};
SkipNode* createNode(int _score, int _value, int _level) {
    SkipNode *p = (SkipNode*)malloc(sizeof(*p) + sizeof(SkipNode::Level) * _level);
    p->level = _level;
    p->score = _score;
    p->value = _value;
    for(int i=0;i<_level;++i){
        p->levels[i].span = 0;
        p->levels[i].next = NULL;
    }
    return p;
}
SkipList* createList() {
    SkipList *p = (SkipList*)malloc(sizeof(SkipList));
    p->head = createNode(0, 0, ZSKIPLIST_MAXLEVEL);
    p->tail = NULL;
    p->level = 0;
    p->length = 0;
    return p;
}

// 输出
void output(SkipList *list) {
    SkipNode *node = list->head->levels[0].next;
    while(node){
        printf("%d %d\n", node->score, node->value);
        node = node->levels[0].next;
    }
}

// insert
void insert(SkipList *list, int score, int value){  
    SkipNode *update[ZSKIPLIST_MAXLEVEL];
    SkipNode *node = list->head;
    for(int i=list->level - 1;i>=0;--i){
        while(node->levels[i].next && node->levels[i].next->score < score){
            node = node->levels[i].next;
        }
        update[i] = node;
    }
    int level = zslRandomLevel();
    if(list->level < level){
        for(int i=list->level;i<level;++i){
            update[i] = list->head;
        }
    }
    // add
    SkipNode *newNode = createNode(score, value, level);
    list->level = level;
    for(int i=list->level-1;i>=0;--i){
        SkipNode *node = update[i];
        newNode->levels[i].next = node->levels[i].next;
        node->levels[i].next = newNode;
    }
}
int main() {
    SkipList *l = createList();
    insert(l, 2, 20);
    insert(l, 1, 10);
    insert(l, 3, 30);
    insert(l, 4, 40);
    insert(l, 0, 40);
    output(l);
    return 0;
}
