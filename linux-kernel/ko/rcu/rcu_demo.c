#include <linux/module.h>
#include <linux/init.h>
#include <linux/slab.h>
#include <linux/kthread.h>
#include <linux/rcupdate.h>
#include <linux/delay.h>
#include <linux/string.h>

struct foo {
 int a;
 int b;
 int c;
 struct rcu_head rcu;
 struct list_head list;
};
static struct foo *test = NULL;

LIST_HEAD(g_rcu_list);

struct task_struct *rcu_reader;
struct task_struct *rcu_updater;

static int rcu_reader_list(void *data)
{
 struct foo *p = NULL;
 int cnt = 10;

 while (cnt-->=0) {
  msleep(1000);
  rcu_read_lock();
  list_for_each_entry_rcu(p, &g_rcu_list, list) {
   pr_info("%s: a = %d, b = %d, c = %d\n",
     __func__, p->a, p->b, p->c);
  }
  rcu_read_unlock();
 }

 return 0;
}

static int rcu_updater_list(void *data)
{
 int cnt = 10;
 int value = 1000;
 while (cnt-->=0) {
  msleep(1000);
  struct foo *p = list_first_or_null_rcu(&g_rcu_list, struct foo, list);
  struct foo *q = (struct foo *)kzalloc(sizeof(struct foo), GFP_KERNEL);
  *q = *p;
  q->a = value;
  q->b = value + 1;
  q->c = value + 2;

  list_replace_rcu(&p->list, &q->list);

  pr_info("%s: a = %d, b = %d, c = %d\n",
    __func__, q->a, q->b, q->c);
  synchronize_rcu();
  kfree(p);
  value++; 
 }

 return 0;
}

static int rcu_test_init(void)
{
 struct foo *p;
 rcu_reader = kthread_run(rcu_reader_list, NULL, "rcu_reader_list");
 rcu_updater = kthread_run(rcu_updater_list, NULL, "rcu_updater_list"); 
 test = (struct foo *)kzalloc(sizeof(struct foo), GFP_KERNEL);

 p = (struct foo *)kzalloc(sizeof(struct foo), GFP_KERNEL);
 list_add_rcu(&p->list, &g_rcu_list);

 return 0;
}

static void rcu_test_exit(void)
{ 
 pr_info("test-end\n");
 kfree(test);
}

module_init(rcu_test_init);
module_exit(rcu_test_exit);

MODULE_AUTHOR("test");
MODULE_LICENSE("GPL");
