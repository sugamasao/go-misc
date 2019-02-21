# subcommand

サブコマンドを使うには`flag.NewFlagSet`を使う。
今回はサブコマンド以下のオプションは同一という形で実装したが、もし個別にオプションが異なる場合は`parseSubCommandOptions`を個別に実装し、`NewFlagSet`を `flag.NewFlagSet(arguments[1], flag.ContinueOnError)` にして、すべてのサブコマンドをパース、errorが発生しないパターンを採用してすべてのケースでerrorであればsubcommandのusageを表示する、って感じかな