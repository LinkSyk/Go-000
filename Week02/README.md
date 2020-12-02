## 学习笔记

这一周主要学习如何在 go 中处理错误，主要是下面几点：

1. 如何让你的代码不充斥着各种 ```if err != nil {} ``` 写法，通过使用标准库提供的各种封装函数简化处理、优化自己的代码逻辑，不需要做判断的地方可以直返回避免啰嗦写法。

2. 如何错误日志和日志上抛更加合理。我们经常出现的问题是一个错误会出现几处甚至十几处的日志打印，既会给程序带来压力还会扰乱我们排错的速度。所以我们要只处理日志一次，如果处理过了，就应该向上返回正常的结果，在底层通过 ``` errors.Wrapf 和 fmt.Errorf("....: %w", err) ``` 等方法将日志包裹住，然后上抛到上层，通过一定的打印方法我们甚至可以知道整个报错的调用栈。

## 练习笔记

个人理解是当 dao 返回sql.ErrNoRows时，应该根据相应的场景来做选择，比如下面两种情况：

1. 根据 id 查询商品列表，如果 sql 查返回的是 sql.ErrNoRows 时，不应该报错，应该自己吞下操作，然返回一个空的列表

```go

// dao 层
type Dao struct {
    ...
}

func (d *Dao) QueryProductList(productID int) (products []*Product, err error) {
    query := "select * from products where product_id=?"
    err = sqlx.Exec(&products, query, productID)
    if err != nil {
        err = fmt.Errorf("query product list failed: %w", err)
    }
    return
}

// service 层

type Service struct {
    dao *Dao
}

func (s *service) QueryProductList(productID int) (products []*Product, err error) {
    products, err = s.dao.QueryProductList(productID)
    // 如果根错误是 sql.ErrNoRows。吞下错误，返回空列表
    if errors.Is(err, sql.ErrNoRows) {
        err = nil
        product = []*Product{}
    }
    return
}


```

2. 根据 id 查询用户，如果没查询到，应该直接报错抛到上层

```go

// Dao 层
type Dao struct {
    ...
}

func (d *Dao) QueryUser(userID int) (user *User, err error) {
    user = &User{}
    query := "select * from users where user_id=?"
    err = sql.Select(user, query, userID)
    if err != nil {
        err = fmt.Errorf("query user failed: %w", err)
    }
    return
}

// Service 层
type Service struct {
    dao *Dao
}

// Service 层是否需要关心 具体的错误是什么吗？
// 直接使用  if err := service.QueryUser(userID);  err !=nil {...} 来将错误当成 opaque error 处理吗？
func (s *Service) QueryUser(userID int) (user *User, err error) {
    return s.dao.QueryUser(userID)
}

```

