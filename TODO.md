# common

- [ ] Настроить git-workflow(автозапуск тестов, бейдж покрытия тестами)
- [ ] Выпустить первую версию v1.0.0


# semantic

## EstimatePayloadOf
- [ ] Попробовать превысчитать размеры для кейсов где используется unsafe и сравнить разницу

```bash
go test -bench=ArrayValue -benchmem > old.txt
go test -bench=ArrayPointer -benchmem > new.txt
benchstat old.txt new.txt
```