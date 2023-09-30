ots
===

A very simple _OpenTimestamps_ command line interface.

```
go get github.com/fiatjaf/ots
```

## Usage

### Creating a timestamp

```
>>> ots stamp file
- stamped digest 039058c6f2c0cb492c533b0a4d14ef77cc0f78abccced5287d84a1a2011cfb81 at calendar https://alice.btc.calendar.opentimestamps.org/
saved file 'file.ots'
```

### Upgrading a timestamp

```
>>> ots upgrade file.ots
- upgraded sequence on https://alice.btc.calendar.opentimestamps.org to bitcoin block 808912
renamed file.ots to file.ots.bak
saved new file file.ots
```

### Reading an `.ots` file:

```
>>> ots info file.ots
file digest: 039058c6f2c0cb492c533b0a4d14ef77cc0f78abccced5287d84a1a2011cfb81
hashed with: sha256
instruction sequences:
~>
  append 6b4276f865eeb47c2be7448494b3c3f3
  sha256
  prepend 650dfc0b
  append 4f448f94f2e708a9
  pending(https://bob.btc.calendar.opentimestamps.org)
~>
  append 6b4276f865eeb47c2be7448494b3c3f3
  sha256
  prepend 650dfc0b
  append 4f448f94f2e708a9
  sha256
  append d1eb4121e8b2fe1f9f5291eca8d0e06a0a857366bafa8f1ad3e1f432aa79e6ee
  sha256
  append 3dae5934c1f58dd5e738d868e6f7e3d42873e01edb46f78ba61afac0a2203344
  sha256
  append bbfa22586eb359091a38133439b56d438dad1a65e2157e072ba748d85d5ad755
  sha256
  prepend 9c95b6b2c47774abbde8e025a59d579fc9dcaaade6655ebf00ad881f11c1600b
  sha256
  prepend daebec6b4c66ca6eb678ca1558416afe70be7cb0ccf5a36c1d493b68f984a5dc
  sha256
  prepend 463975ff43cfd25e84a14b7f908d6db37bc6abac9434188f9a56adf6b9a55d71
  sha256
  prepend 97dcb35a516d2db80554b0f0b79d568580a0731e59f702da83b74afd68578d8a
  sha256
  prepend 249a7742256a87c0f41ad1907386abca4dd3174c5b7322214fe7cccfd94a17bf
  sha256
  prepend ea1c733f3d9c63b1e0a86a91a205d4e211634e02ac007457c135ada6d8465136
  sha256
  append 190d0cbfac7324e4c8decb34a837c4f5db1eda7989ad5787a698df5f1475fcdc
  sha256
  append 39625a35caae7abd3cce58de18ea4f029e8ad7083ecda4ad0fd15729d5108b99
  sha256
  append fbb6eeb1592e1feeb96f85ac794d059948731797dd44ca052dd55c2d52dc50f4
  sha256
  prepend f7f2d9871f28d8c43b6b0b2b6106c073daa3d501485d4aa5a2b0dac276e78dd1
  sha256
  prepend 9c2e05de8ce8a21d0184c03a083999410799187f8885665686852f895e878d5d
  sha256
  append e730a41c7751a1b0763704b8ec7a9f5bcca14b08a1b00304081a0c04eb05677d
  sha256
  prepend 010000000100d575ce1d81cc704aedd9deb12cc806651d1f00d5ce7fb3b090513b1ea538f20000000000feffffff0292b90300000000001600147fc8a2fcea10426
6fc71727ddaca64488f64d7e70000000000000000226a20
  append cf570c00
  sha256
  sha256
  append 8cc9ae705840ad8e49453eb0af426a2dd62a8b140d2358aec73f9b7973b62866
  sha256
  sha256
  prepend 72a1950b7fcb24a609bdff45c16fd45406252dea43a4baa63f8a837f32458b81
  sha256
  sha256
  append 7a814af76526d43faa409b074a77b9ffe0db743576cf6163eea5282c4be8019f
  sha256
  sha256
  prepend 989de4b6bb30be091e2020558275f30b84d889fa1d551d2ce7d4c636a62bad94
  sha256
  sha256
  append bc2eb0d9c635e41218551c827b6d39a043062e1c94c9a5fbbb83587afe5da2fa
  sha256
  sha256
  prepend 9391adcb3de07623b0f94036e3302ebd3cd2b36c77d7df0e8e3b51f97fdf7ae0
  sha256
  sha256
  prepend e11d1d75fc63f65a5f852334124675f0321dfd5c2ef7473b81182c4fc05336e9
  sha256
  sha256
  append c09fc89327c0882d6c0605ac6baacea39cbcb03144973dad7d49f73a8586f19a
  sha256
  sha256
  append a5f00d9d78a2b7af720230345bf3a28659bb383c1b0b243e34802307fcf43d9f
  sha256
  sha256
  append 385b78d6a0a5fb15335dba751a7c44f3bae1a1201f7501223532f438661553eb
  sha256
  sha256
  append 4e51b160eda86c048ca9b6987c1fbad40f9d0bdbb9ae8a245a7cd648e7a02b53
  sha256
  sha256
  append 73dc3ed5877addda78d4fffc99783ae7adea4e6e0f7ade5f8df6841e8bfa2acd
  sha256
  sha256
  bitcoin(808912)
```

### Verifying a timestamp

```
>>> ots verify file.ots
> using a an esplora server at https://blockstream.info/api
- sequence ending on block 808912 is valid
timestamp validated at block [808912]
```

## License

Public domain.
