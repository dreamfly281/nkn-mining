<template>
    <div class="nkn-setup-page nkn-card-shadow nkn-after-clear">
        <div class="nkn-create-account-panel">
            <label class="nkn-page-title-label">setup - step 3</label>
            <h1 class="nkn-page-title text-main-blue">Download wallet</h1>
            <div class="wallet-info">
                <div class="wallet-private-key">
                    <h4 class="info-title text-main-blue">Private key</h4>
                    <div class="info-text text-label-gray info-private-key">{{privateKey}}</div>
                </div>
                <div class="wallet-address">
                    <h4 class="info-title text-main-blue">Wallet address</h4>
                    <div class="info-text text-label-gray info-wallet-address">{{address}}</div>
                </div>
            </div>
            <div class="setup-button nkn-after-clear">
                <button class="nkn-normal-btn" type="button" @click="signIn">Sign in</button>
                <button class="nkn-normal-btn" type="button" @click="downloadWallet">Download wallet</button>
            </div>
        </div>
        <div class="nkn-setup-page-wallpaper">
            <img class="nkn-wall-background" :src="'./static/img/wallpaper.png'"/>
            <img class="nkn-wall-pad" :src="'./static/img/wallpad.png'"/>
        </div>
    </div>
</template>

<script>
  import FileSaver from 'file-saver'

  export default {
        name: "ns-show-wallet",
    data: function () {
      return {
        wallet: null,
        privateKey: "",
        address: ""
      }
    },
    methods: {
      signIn() {
        this.$router.push("/main")
      },

      downloadWallet() {
        if(!this.wallet) {
          return
        }
        let walletData = new Blob([this.wallet.toJSON()], {type: "text/plain;charset=utf-8"});
        FileSaver.saveAs(walletData, "wallet.dat");
      },
    },
    mounted() {
      this.wallet = this.$store.state.global.wallet
      if(!this.wallet) {
        return
      }
      this.privateKey = this.wallet.getPrivateKey()
      this.address = this.wallet.address
    }
  }
</script>

<style scoped>
    .wallet-info > div {
        padding-bottom: 40px;
    }

    .wallet-info .info-text {
        font-size: 14px;
    }
</style>