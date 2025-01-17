package main

import (
	"math/big"
	"perf-test/spammer"
)

var (
	url     = "http://3.79.42.126:10001"
	chainID = big.NewInt(8787) // unifi testnet
	maxTxsPerAccount = uint64(1)

)


var (
	alice   = spammer.Account{PrivKey: "7ac9a7e636c0f437dae18e388faf8790a96e862c2f1ca517dc671e7d1cfe0348"} // 0x3ad4C1808d207433d6a476411044Eb6eE38b4849
)



var (
    a101 = spammer.Account{PrivKey: "528af91dfe6951b504b4209973eb3e6040f50decb0fd49623bdd78f0e62de2ad"} // "0xE02FE605f08c43E46167068561c79aCa5cFA9543"
    a102 = spammer.Account{PrivKey: "55b0fc33494d2ac7905b89e5525041b5648daded7f1c85496d98ca1ed2483eed"} // "0x8DA824c71878c2e8061494ff3fF720A83Fc9A3E4"
    a103 = spammer.Account{PrivKey: "1ddf81f3f6cbcec3ccd1123f3ae98ab276f3ddae40ae6cd1848ef71b6ff78a4d"} // "0x5E9b6aEb6d3B35E0D021210Bf2ECFAb1F1399bff"
    a104 = spammer.Account{PrivKey: "ac26a4fafe7add9664d9d84d62110522cd9bce10255bd9a456d9e549adba1b27"} // "0xdE4c9fa2b7f80A7F661f3Ca50C7869A6A30a3dA9"
    a105 = spammer.Account{PrivKey: "bad05810e864a6f2d80e126199480fd414c55fe5096d7a1e59f225968f5cbbd8"} // "0xc4C664d6d5eCA25453696F35E27C87fF9090652a"
    a106 = spammer.Account{PrivKey: "2a5c7b3f1fae704c055fc0c084b0d9c772d026acb986b1ef3baf366f142bbf90"} // "0xCA3Ef4c0Fa6dA5176E6B9605E768D08fd1d5434b"
    a107 = spammer.Account{PrivKey: "0a6b1d084a57e0c51043c8d272753c1b364d51bb1a7b71572e41a8554c98915e"} // "0x2E1ab310F3c8d51B2b7a4917518616cd2409f1DF"
    a108 = spammer.Account{PrivKey: "2f81a9b81c060c3ad6c9ffc370b191f30e184391316fa0eb87c456312d8e4e5e"} // "0x46E52E4623FE854e755a3178568721908ac514B2"
    a109 = spammer.Account{PrivKey: "cfebd5f0a3b1cc25134f0df9798a213fd4a52f3cf9de78082753d5e67277433b"} // "0x17a25DeB1124Ec0415656bF7AB9ACF6472Ec4b0d"
    a110 = spammer.Account{PrivKey: "f8558241793cda1b0e09ffb841102b2f93dc6778a8950862373b22cbcb88254d"} // "0xfc3735FCcbc2d3F30e42B9C8169B0aCC86B0B366"
    a111 = spammer.Account{PrivKey: "f16bbb66915fa5d431c6231f40162ca4e4348f75645a72c05d1dcaa762b24335"} // "0x004B98a06862bAdAfDEe6bd5D88009b3f45edDF5"
    a112 = spammer.Account{PrivKey: "ac3316326a2543a3f1f968681c612bd06bfee6c4b4daa72234b67d1ca89d13aa"} // "0x53dD7a0fb17Aa800f345c41967694753200ac841"
    a113 = spammer.Account{PrivKey: "40f1109a74484e0a3cdc89e9bc92b3ce45fadfafd074761ab45ad9db02cb0f17"} // "0xC3FaBDd27FB7957ED8600e9A19850A9194c7E428"
    a114 = spammer.Account{PrivKey: "643a8f792f7124daab74d36f9706d5c69c9d6b2399fdd7c59d9c065d3bd9f4e8"} // "0xCF686f16E84CfFB75d39ADB4c6f5A4c708EC36D4"
    a115 = spammer.Account{PrivKey: "2563115e405e0c84aae299916a3898fda377b46f819acc77c48d6546a7db01c4"} // "0x1aCD79190fb140dD3983b19deB1CdF798A72d43b"
    a116 = spammer.Account{PrivKey: "ed8025e0507bd631afe249f93e65080f99ae9330db13117895f448c102979796"} // "0xBa0a1026b3A82389c468D9d70dcd54c2Df2A2dd2"
    a117 = spammer.Account{PrivKey: "282f6e52a4e69c18b861a1ded171e5b1f62a50c1c2e5b415ee9a91ecfe96a339"} // "0x3d28B769343bDf3fE70703C40407d6A042D1fB62"
    a118 = spammer.Account{PrivKey: "25a5a5c0c3eab6d45ca079353cc4b92fccabc8c30cfd6706c4f56d647dedf53a"} // "0xaB2bf01ad015b3dDD6A72bAA8F6b1e78EF7c9E55"
    a119 = spammer.Account{PrivKey: "94ea137ef32208da3e7e4017bfed3d92a6a169484d6c7bc8391e22586f0766a5"} // "0x4Cd4DB7C332D3992A56dF2C9b368E75e590Cd76B"
    a120 = spammer.Account{PrivKey: "32d1915536aab3d7280b24a5231f1488f2c75da1816a2e15f97092915f60e18b"} // "0x22e3C5Bc697E6618676D7eA97120F0334ebEAcdA"
    a121 = spammer.Account{PrivKey: "77de8219d5ea124b7b58b6f5ce5a5b3a36123a07925857a4fd5fff1ed2223df3"} // "0xA860094dfCFf5c1B44925ff92C63de98ea58ef3F"
    a122 = spammer.Account{PrivKey: "efb02b8ae438cb07d4d08e7d5d16c2b3580aea48110c075fed78821524842d54"} // "0x94E9CE7e6A54740DC60EBAD56014b6D2BD7adfC7"
    a123 = spammer.Account{PrivKey: "3f24f66a8dc5898dcd38c7b20fc4dfa6e37ba447c420e1641764e45b73e0d7eb"} // "0x085B93E0342653FEAAc671aDE6AB9c679A55a236"
    a124 = spammer.Account{PrivKey: "0d34af229254d6f235b33def8981b4874114bb591ab8ed3d7ff61f290020ddff"} // "0xefF86B2BAcf70162be52Ed43720a041004A4a593"
    a125 = spammer.Account{PrivKey: "225f7e5fdd3683b012c561261e8495624bbd83b83ee394267787e47991ae39f8"} // "0xAe1cBAb01e5fea97355Fa73D3551C3f956059cB3"
    a126 = spammer.Account{PrivKey: "c664f0b9e3ac2dadfb3ca25d65fec990362a2a183d7e3e377b6a3b3de92a4e96"} // "0xc7B9F223E9c5bF4d973f52630B28FBdB4Fe4cd22"
    a127 = spammer.Account{PrivKey: "123881d875f57a09d08dfd99482dc64908cd7c9de8c79a2f3be04a57087b33c4"} // "0x37B302098501795674E2B718b33E5b0E3Ee138c7"
    a128 = spammer.Account{PrivKey: "72d158bf7714925f41f43b3c98c756db3bbd1eb73a0d310e1674db4f2a8636ce"} // "0x21A470a827237ACeb7d738C647b982e4EABBd043"
    a129 = spammer.Account{PrivKey: "8a7ee51d987c826a96a327405888d963b8d7b09cd2d4dfd828c1e54073194a40"} // "0x5335648b931B25DFAb00A6A0191c6fA6048d3BcF"
    a130 = spammer.Account{PrivKey: "02d3cc39105c31b2ac161688850684da716a15e15cc02c04c3230abea70d0119"} // "0xc80efCaA06564fe4e8a98a3DA39230f7F6Af6695"
    a131 = spammer.Account{PrivKey: "cd3bf9067c637a54edc7e2826fea380d52dd7fc66065a2f73adb122614e6d581"} // "0xcae66FB8464fE703eb72e7CD4A7F19fFd4fa3c2B"
    a132 = spammer.Account{PrivKey: "8a5729ad6c6d2adbfa0e1483fd682b5a2828e8d45ffbf1f401e56e3a225a1ff7"} // "0x492EB5C60a25Bd5A05C20871e46b8EEa4d4A2324"
    a133 = spammer.Account{PrivKey: "a92e8863841743dfd9c30f0021d07b02cccb7c3cefaf01f74fbdcb3ea3dceecb"} // "0xc7d93ED387B12572edC043296a7FB20ba8f8Bb74"
    a134 = spammer.Account{PrivKey: "c34d14cb9978a197d63d1b4f7c8613a459f00c123f2b500f255e4b21c4cea78a"} // "0xe5e90410ff32a8D223B9153b6BCe2507Dd372976"
    a135 = spammer.Account{PrivKey: "78c8451d3e0fbb234d4d3feda394cbe7ba645cc5d53ad9cdeee9369d3e2dc3ca"} // "0xa0d48Ca5e3378AbE8B06B10741aCc85C5269Cd31"
    a136 = spammer.Account{PrivKey: "1ab469d9fc217cf82f8ad700e99b38ca09e587f843626e38e1d269b9384d392e"} // "0xe4E9f803b828EfA84ea24E94a6d4CDcC2a791d80"
    a137 = spammer.Account{PrivKey: "e3c731b9b6f977243aea40d3d8fca43f1b6917a7dfcc96aca182c02d23c088f0"} // "0x0869c342308828e69D992fB5F332c5503146FdB7"
    a138 = spammer.Account{PrivKey: "c6b05991716cf87f3684b31a32f734ef306507c3f5b3cc769f1be55026d5d881"} // "0xd87A7f262BaEEbA067B5CF64560C89014aC78a6F"
    a139 = spammer.Account{PrivKey: "29a72d9df87ac5d9f8bef0d3540810d50532a5b4800d220456a8427816e7ab2d"} // "0xb85b8E06af1de6fdD4c04b67C1fCC9ad5973BF41"
    a140 = spammer.Account{PrivKey: "a0522ebdad4466b3e41abd15657a13d96b114614b42e503cf10d9ad57139bfa7"} // "0x348371723fFD72fCDB0435dB4BE10B1fdb3c46AD"
    a141 = spammer.Account{PrivKey: "7252b5b405cb7213ca5fbb3374f4c3f12ffda8b715abf836031fce589dc7dfaf"} // "0x0c8D582595e5E8C68b6912A5fA2150663fa2A6A3"
    a142 = spammer.Account{PrivKey: "feb7420dbb39091888f2cdab4dda01a2791d6b959fbe15ba163067e86a442bf9"} // "0x83b339C39Ee07BF8bB78b71dA7d3F2Eb1a1415c3"
    a143 = spammer.Account{PrivKey: "52aab43ab0153bdc58d5799d23e983e04d3d25524f61cbf078a33a41c8a54be3"} // "0x26A2d3977dc8B4e803C27523E315cA74D662f9d8"
    a144 = spammer.Account{PrivKey: "99eb65afff5afe8b52a37c470189e3ebeff747989ab876249289dd42d0981550"} // "0xA6b68b85894CC23d23bA99B9a0B5AA1c0E09e558"
    a145 = spammer.Account{PrivKey: "30a8badb9e5e2d239bb888d99169df50838c0f2019d21bc453ece960e65d36ee"} // "0x5160DEeB8E5510852D886e5CF74282C8C07b9e0B"
    a146 = spammer.Account{PrivKey: "048743355155c8dbf15726c4972bcb4666f12385ad48eed615d3df18f6b64351"} // "0x6Cd22639757B2c8f9461E44e0607e0DCfa3a3760"
    a147 = spammer.Account{PrivKey: "a484bd9e5d123e01d1736f203ff8d0ebdc52851657fdc0120e8695c0fc94d4ff"} // "0x968165e422E28693450963CFF9f2DC19be56b428"
    a148 = spammer.Account{PrivKey: "eddddf26e790c567a40b9a001ee187c96c3d6213f2a6af4dc3126ae7459989e8"} // "0x0a5539cAB1f2618Caba195F9d416ADD47747eA90"
    a149 = spammer.Account{PrivKey: "e6f476f6509f38404326b581f339cf1bf51b574c75cf5dc393ad0baa5e860767"} // "0xD0D2aeeCd3d13d47F24056F7729a7e1010c57563"
    a150 = spammer.Account{PrivKey: "dba8d14280c80792a24b3655bc1655766f6237bb1f3b52c8179bf3febf290e78"} // "0xD5DF4989b3cb9c917c338856a6C9Cdfb8b8c9179"
    a151 = spammer.Account{PrivKey: "50e3846eafe5fff8cb8305a5038abf3f1ba788b858d5705d059311bc9be8be08"} // "0xE29bC56480Fe24Ca47Cf4F7285929AF3D6EC57Ee"
    a152 = spammer.Account{PrivKey: "a0de8dd345946889149a8c50d215813d9b33c3ac5b13759ede97dec831a424c6"} // "0x242bE95e97B03Ef702a458597C721dd6f885CFb1"
    a153 = spammer.Account{PrivKey: "e1394ccb701c9a55c58bc0d5628ab947e3e6afa92b5c4b94a6893d9795434a03"} // "0xBE16453d9e4cb18a16FEA098bb5C5355aDBB9129"
    a154 = spammer.Account{PrivKey: "9ed6a58f355b6abda187ce3c587758f0f6b450f05c5e69e04ee7a0b21355cf6b"} // "0xb293B3590c10Ea2B6B0676ffe33E1e46e00bFDE3"
    a155 = spammer.Account{PrivKey: "e18a6795fb2123056ca13c05946eba2f5d683ddc7f6b3c8fe3b00310feabc340"} // "0xACf762e3752AA2ba2bF42649066Ab8dae8D4DCFD"
    a156 = spammer.Account{PrivKey: "a3289de174583f35cd5afffc67bcb8344cd4e9a4b60255fadc6c888d15e2ba63"} // "0xC667f3EC827FCbc21F0007C8922AECcE87664c94"
    a157 = spammer.Account{PrivKey: "0454df308bf65c3c2b8ac1a5194c9590952637bf672fe044d06f92d8aecb0147"} // "0xC6775ff92525Bf195b010E8833e4f22F5fd2e516"
    a158 = spammer.Account{PrivKey: "a0c4a8d1f1e5dd65fe82ecc8ea9fa933c9fb13ce861a659bef3205eab03cecd8"} // "0xF6e451bE2C2f075444489437c1276E7b3FB72f94"
    a159 = spammer.Account{PrivKey: "9bc2b2110ebf22a23c394af1046307167a82c18fe6de01f594023abf1d83e143"} // "0xC88c3c82Ba858311e99a519a7D7EBBa71c1Bd3a2"
    a160 = spammer.Account{PrivKey: "18c70825f60f1e6f7b40706461eca66e9bf6752b08d03853d69cbcc36c04195a"} // "0xCbAFc93F4E2425b495C6850d181cFf72d4E8baEE"
    a161 = spammer.Account{PrivKey: "40dc8af78098cb6746bbab4206660636d118fdcb2c9d2fd4c982e2d11c9b803e"} // "0x36367bf8D0A6dDE8d2107Afda8f84EAC9601FaE4"
    a162 = spammer.Account{PrivKey: "06fbc4128a0e424cbc38d2cc63af216b17de0c8c26c761237d7bf4dc488860fb"} // "0xac83D1A30B99277f6424f0D8cB677419bC232953"
    a163 = spammer.Account{PrivKey: "4aa9856848ef77eb6d66a97bb0805f3b904d01a7597dbc517da87473647373be"} // "0xBb490276E9a1A0E85EA18Bf44e46688631C2b6Ee"
    a164 = spammer.Account{PrivKey: "4956ee2b53b962327718817620c0dee4fa788939b61218c62380ce9220331bbc"} // "0x753B24e0eE6B5EeC05af7d4de7890E55701cB70a"
    a165 = spammer.Account{PrivKey: "d34881d74eeb538a3dcd82a8ce589d6711b0fd76d924739e0969d716c0db342b"} // "0xa012AAc35da87a04E9f7dd26247AC12f73c86b3a"
    a166 = spammer.Account{PrivKey: "fb03f8395f32b4909fa79cd6e13b02691c802c630dcb794a440b65cb3d14c54b"} // "0xd77dc1bb6c115e7981Be8CC7D61A3e11983c6342"
    a167 = spammer.Account{PrivKey: "4927860cb96d6d869a2c2d07ee818f0008a5b2c71d193325b8f6d50ab58c59a3"} // "0x21998a5c75Ca955C8AFEB2c5DF1AC239eE300c12"
    a168 = spammer.Account{PrivKey: "c2957eaf56e15dc0c153f4fb04bbb466d643671426930f26fd43f06f49973a1b"} // "0xCa17F55ac5F88c1e3F381428c052303CAcC80C76"
    a169 = spammer.Account{PrivKey: "c1080b1af4d98e49cb6a31abc9e7d5e95b1361ab9c6bc389e6ae44b3e73cc3b7"} // "0xD37814E4a8143A24c8901DBDdC9E727D03960CfD"
    a170 = spammer.Account{PrivKey: "235ff8a1b1d7e02e09107052d69b670829a6998725fa0b0697534fc9d4a0b115"} // "0x749449231bd6B4c2A34e59474920862331997B4B"
    a171 = spammer.Account{PrivKey: "cbc0419c14b50c15825cb5afdcec24e6f454aba733d037d5760b1e64c97fa01d"} // "0xffF97E27af1430c5aA1543Ba5Cf4adF51A621c5d"
    a172 = spammer.Account{PrivKey: "cde061e9ea3f3790fd4c4ddef6b6ca4f96d79ecd4b4b092fda251f9a4a77e637"} // "0x15d6093D974A42C3921bbC3dB63c283275aD3C84"
    a173 = spammer.Account{PrivKey: "39d74b4377e5d0ebb2f0109b5c88a328208bc10fdb337b3e50436e3382561de5"} // "0xF11FA374cC7B1Ed367a4Aa3AA8db1e444F227a90"
    a174 = spammer.Account{PrivKey: "da122ccc69b2b4265e68987a086b386a295749f2cddc4f43f4fc28f9a42ca19e"} // "0x8d45E384CFe3dDE8a2559F3471cB5812110EFff0"
    a175 = spammer.Account{PrivKey: "08a7ca779282bee7e4367365ab4fdad4bedcd71fa1a6ab446dfdf12265fb0e99"} // "0x5FBea2A29524bd1b264a8b1aeDd0ca7544F8aD52"
    a176 = spammer.Account{PrivKey: "4fb22460e5a56b89eb31243ff66fb3206742f308e5959f84922482089d00fe42"} // "0xeC1db2Bc81e8cA036D0e2bF96f3CBd59f858855e"
    a177 = spammer.Account{PrivKey: "8d8916c91f5e3cb8c73885aa27559eb9c1462de6036a9c32ec90b1cdc2d38e88"} // "0x921A3E8F270D162e02F12F79ad4e15EdbEb82afF"
    a178 = spammer.Account{PrivKey: "b112880110086874e40afe944670ac1fab62edd6152e72edc4ab8dc8e0d2892a"} // "0x4501ADc040ABC7C1650E65c3a2656b9c34e45D9a"
    a179 = spammer.Account{PrivKey: "f684321f88a7e09502d55f10b4bb40e0878910caacf9a6d00bd4125e766807d7"} // "0x57F080eA3FF11D6e8e65C225d3B728Bb3866b564"
    a180 = spammer.Account{PrivKey: "993633287dbf01130a5e4e176d6e681598b8e4d656465e4e380d668c6da76f1f"} // "0x39d074d8a9292Ec51CBb385bbB906E8bc48A59d6"
    a181 = spammer.Account{PrivKey: "f7539587a4044ce5fb56cca1c67c6c65ab27d68fafedb068c9995f99aeea8a0a"} // "0xD3aCda3629F559031190a5F2A91a93760E7B42A0"
    a182 = spammer.Account{PrivKey: "816e5ffe17c1434782ee0ba89e3914a36086a96a69ad166842cd4095a1c31dc4"} // "0x0903d4E3Bb1358ef0eF4085755AFaC0fa17e4e27"
    a183 = spammer.Account{PrivKey: "60703f843cb7601e4ec53e87dc10a4ccee321536da56e7e6c698a83ead841e02"} // "0xeA5B900bc16027ae447eFEA1E1617f7c84243Ca4"
    a184 = spammer.Account{PrivKey: "fb225c85ebb22027fb8abee73b1bdd45e06350141a0afa645fceefe48b18f82d"} // "0x686A6d8a57aDa4ec35832C048B1227C0C838421B"
    a185 = spammer.Account{PrivKey: "06a6db8f7e99c32d6657308eaccbb849a78374606c392dafb2024df6cfac02c8"} // "0x2c9eab958F057143DE6575802E6Aa6b9C1022857"
    a186 = spammer.Account{PrivKey: "047a51df082753ff19e8a0d3506decca9dc1b500d3445a1e0604691e05c310ef"} // "0x81DeB3a8D66a977B87b3D91083C18366D28453F3"
    a187 = spammer.Account{PrivKey: "58a7b0a8a4b6bdc6894b140a2696359a1c06791f40e53e224cf1d837522fec7d"} // "0xcB739DD565037fE1fe59C866766A18dc5bAb1e3C"
    a188 = spammer.Account{PrivKey: "090863b0258ca9ba2f518ba9398d7b4f5fa1eae3c82ccc19ec3db7edf3e3e8e8"} // "0xcB0b657252fefF4cc34563290AEee0Cb44d2c470"
    a189 = spammer.Account{PrivKey: "d5e50c5988841713a00c752514f2927e3deaffc743ee9096c5d20365b43ddd46"} // "0x27E6Df7C3F6c0277741138af10da114F10535ea7"
    a190 = spammer.Account{PrivKey: "c943ab1d2f64d64ed0b95ae179ec6ca55976749fcfb01ab2ccab28359d6e2f51"} // "0x61fCB3eC88D8204bEEf04bF30375b7004A31170F"
    a191 = spammer.Account{PrivKey: "060bf8f398e146d0192f64f35f9f8c24b018f4221a84b3c6fc0f89771c314d81"} // "0xd146F4E1F5d59a28a39083bA55CB652aF42C32a2"
    a192 = spammer.Account{PrivKey: "52e41d4ba5bfc2acbf2ec4a2975842cb24c7fb2034f064597e6155c7ec7d6b56"} // "0x5E2f889C7514B2e035d821bb269859ed1fE0Bdd2"
    a193 = spammer.Account{PrivKey: "7db66a2bf4271932949d24e684fa710c6a75d38a45fa12e73525c50fdbe09bf2"} // "0x7C7D78fC9641ad28E081DddfeF419eC4D722f079"
    a194 = spammer.Account{PrivKey: "0beb19c806bdcbde7b3ebb1d726ea941f399323b35a0ff014fe371d1cb467488"} // "0xDf5570e42bA8F3a203a39a4D4f8Ef5b0B73fda99"
    a195 = spammer.Account{PrivKey: "7a381cf686577d6bdab9072b7542f665efa4d308ffad3a10cafa8d3cd4113cbd"} // "0x54ef5D6a42D9019ce00B01f139F0599aA5A4277C"
    a196 = spammer.Account{PrivKey: "d9ecc23d368f04a41fc49030958b9ed5baba4733cdd0201c1bb3abddf9a230ea"} // "0x53Da7Fe6BA27a4Dfeed02560e96ccD30A5Dc63cF"
    a197 = spammer.Account{PrivKey: "b6ad5cb8eb9c3d9fce454c9ca986b74f0260d0b07351ecd0d34c623c7b09f078"} // "0x732d8DFcbeE3B9e2B6a42c456Dc6eDeD719A1499"
    a198 = spammer.Account{PrivKey: "1d99ca6dc3b51d045473a5f47afbc1df76f77f95a66779416d25444650cb74f4"} // "0xAD591d2B972d7f1984608878182F5dcb4A02b44F"
    a199 = spammer.Account{PrivKey: "52e41d4ba5bfc2acbf2ec4a2975842cb24c7fb2034f064597e6155c7ec7d6b56"} // "0x5E2f889C7514B2e035d821bb269859ed1fE0Bdd2"
    a200 = spammer.Account{PrivKey: "7db66a2bf4271932949d24e684fa710c6a75d38a45fa12e73525c50fdbe09bf2"} // "0x7C7D78fC9641ad28E081DddfeF419eC4D722f079"
)

var (
	accounts = []*spammer.Account{&alice}
)

var (
	prefundedAccounts = []*spammer.Account{&a101, &a102, &a103, &a104, &a105, &a106, &a107, &a108, &a109, &a110} 
)


