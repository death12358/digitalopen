# SlotMathServer


## Getting started

```bash
go run app.go
```

---

## Dependency

<details>
<summary><b>github.com/bofry/host-fasthttp</b></summary>

  ⠿ _Http Server 套件_

</details>

<details>
<summary><b>github.com/shopspring/decimal</b></summary>

  ⠿ _點數計算相關皆使用 decimal 套件，避免 IEEE 754 精度遺失問題_

- **Install**

    ```bash
    go get -u -v github.com/shopspring/decimal
    ```

</details>

---

## SG001

pickem 為選擇不同 unit bet 遊戲，可選擇項目有 8 18 38 68 88.

```bash
curl --location --request SPIN '127.0.0.1/sg001' \
--header 'Content-Type: application/json' \
--data-raw '{
    "id": "20221129",
    "brand": "test_brand",
    "username": "gory",
    "pickem": [
        "8"
    ],
    "currency": "ts_coin",
    "bet": "8"
}'
```
