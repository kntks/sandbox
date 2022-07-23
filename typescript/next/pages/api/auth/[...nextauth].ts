import NextAuth from "next-auth/next";
import CredentialsProvider from "next-auth/providers/credentials";

// example: https://github.com/nextauthjs/next-auth-typescript-example/blob/main/pages/api/auth/%5B...nextauth%5D.ts
export default NextAuth({
  providers:[
    CredentialsProvider({
      name: 'Credentials',
      credentials: {
        username: { label: "Username", type: "text"},
        password: { label: "Password", type: "password" },
      },
      // example: https://next-auth.js.org/configuration/providers/credentials
      async authorize(credentials, req) {
        console.log("credentials ", credentials)
        return { name: "admin"}
      }
    })
  ],

  /**
   * 開発環境: optional, 本番: must
   * NEXTAUTH_SECRETの環境変数があれば、secretオプションを設定する必要はない
   * https://next-auth.js.org/configuration/options#secret
   */
  // secret: strings, 
  
  /**
   * strategy: "jwt"の場合JWEをcookieに格納する
   * "adapter", "database"
   * https://next-auth.js.org/configuration/options#session
   */
  session: {
    strategy: "jwt",

    // sessionの有効期限
    maxAge: 60,

    // databaseにsessionを更新する。0に設定しているなら常にアップデートする。jwtを設定しているなら無視される
    // updateAge:,
  },

  /**
   * sessionで{strategy:"jwt"}を有効にすると、セッショントークンとして使用できる。
   * デフォルトで暗号化されている
   * https://next-auth.js.org/configuration/options#jwt
   */
  jwt: {
    // デフォルト: session.maxAge
    // maxAge: 
  },
  pages:{},

  //https://next-auth.js.org/configuration/callbacks
  callbacks: {
    async jwt({token, user, account}) {
      console.log("jwt func called", token, user ,account)
      if(account) {
        token.accessToken = account.access_token || "this is access token"
      }
      return token
    },
    async session({ session, token, user}) {
      console.log("session func called", session, token)
      session.accessToken = token.accessToken
      return session
    }
  },
  events: {},
  debug: false,

  /**
   * default: https: true, http: false
   * https://next-auth.js.org/configuration/options#usesecurecookies
   */
  // useSecureCookies: false,


  cookies: {}
})