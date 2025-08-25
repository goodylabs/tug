# TUG (PL)

![tug](https://raw.githubusercontent.com/goodylabs/tug/refs/heads/main/assets/images/tug-cli-logo-256x256.png)

TUG to narzędzie cli które pomaga podglądać rzeczy co się dzieje na stagingu/produkcji etc. dla danego repo, kiedy nie chcesz się bawić w ręczne wklepywanie komend.

## Instalacja

1. Uruchom poniższą komendę, aby pobrać i zainstalować najnowszą binarkę TUG:

```bash
curl -s https://raw.githubusercontent.com/goodylabs/tug/refs/heads/main/scripts/download.sh | bash -s

rc_file=${HOME}/.$(basename "$SHELL")rc
echo 'export PATH="$HOME/.tug/bin:$PATH"' >> $rc_file
source $rc_file
```

2. Zkonfiguruj własne ustawienia w TUG:

```bash
tug configure
```

## Podstawowe komendy

```bash
tug --version
tug --help
tug completion # komenda wygenerowana przez framework - na razie jeszcze nie rozkminiłem jak dokładnie działa.
```

## PM2

Ułatwia uruchamianie komend na zewnętrznych serwerach, jeżeli projekt korzysta z pm2.

Co to robi tak po kolei?

1. Zaczytuje konfiguracje z `ecosystem.config.js` w katalogu głównym repozytorium w którym się **TERAZ ZNAJDUJESZ GDY KORZYSTASZ Z TUG**

> czyli jak chcesz podejrzeć co się dzieje na środowiskach `<project>/<repo>` to musisz wejść na lokalu do repozytorium `<project>/<repo>` i tam uruchomić `tug pm2`.

2. Jak już zaczyta konfigurację to zadaje ci pytania o to co ciebie interesuje, czyli:

   a. jakie środowisko (staging, production)?
   b. jeśli jest tam więcej hostów to który host?
   c. który proces na serwerze?
   d. co chcesz wykonać np. zobaczyć logi, zrestartować proces?

3. Po tym wszystkim generuje odpowiednią komendę i ją wykonuje na zdalnym serwerze.

### Nawigacja

- strzałka w górę/dół - przechodzenie między opcjami
- enter - wybór opcji
- ctrl + c - powrót do poprzedniego widoku

### Przykład użycia

```bash
# Sprawdzenie do których hostów masz dostęp
tug pm2 --check

# Podstawowe użycie
tug pm2
```
