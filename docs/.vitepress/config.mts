import { defineConfig } from 'vitepress'

export default defineConfig({
  base: '/freighter/',
  title: "Freighter",
  description: "A simple but extensible and unopinionated load-balancer written in Go-lang.",
  themeConfig: {
    nav: [
      { text: 'Home', link: '/' },
      { text: 'Guide', link: '/guide/getting-started' },
    ],

    sidebar: [
      {
        text: 'Guide',
        items: [
          { text: 'Getting Started', link: '/guide/getting-started' },
          { text: 'Backends', link: '/guide/backends' },
          {
            text: 'Strategies',
            link: '/guide/strategies',
            items: [
              { text: 'Round-Robin', link: '/guide/strategies/round-robin' },
            ],
          },
        ]
      },
      {
        text: 'Tutorials',
        items: [
          { text: 'Custom Strategies', link: '/tutorials/custom-strategies' },
        ],
      }
    ],

    socialLinks: [
      { icon: 'github', link: 'https://github.com/ShardulNalegave/freighter' }
    ]
  },
  head: [
    [
      'script',
      {
        src: 'https://www.googletagmanager.com/gtag/js?id=G-JCJH9KGFWV',
      },
    ],
    [
      'script',
      {},
      "window.dataLayer = window.dataLayer || [];\nfunction gtag(){dataLayer.push(arguments);}\ngtag('js', new Date());\ngtag('config', 'G-JCJH9KGFWV');",
    ],
  ],
})
