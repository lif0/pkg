# common

- [ ] Настроить git-workflow(автозапуск тестов, бейдж покрытия тестами)
- [ ] Выпустить первую версию v1.0.0

- [ ]Бейджи как тут https://github.com/kshard/chatter/tree/456d9ae64dea7d70b774b857ab8377032509397b?tab=readme-ov-file

# semantic

## EstimatePayloadOf
- [ ] Попробовать превысчитать размеры для кейсов где используется unsafe и сравнить разницу

```bash
go test -bench=ArrayValue -benchmem > old.txt
go test -bench=ArrayPointer -benchmem > new.txt
benchstat old.txt new.txt
```