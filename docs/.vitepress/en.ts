import type { DefaultTheme, LocaleSpecificConfig } from 'vitepress'

export const META_URL = 'https://lancet.go.dev'
export const META_TITLE = 'Lancet'
export const META_DESCRIPTION = 'A powerful util function library of Go'

export const enConfig: LocaleSpecificConfig<DefaultTheme.Config> = {
    description: META_DESCRIPTION,

    head: [
        ['meta', { property: 'og:url', content: META_URL }],
        ['meta', { property: 'og:description', content: META_DESCRIPTION }],
        ['meta', { property: 'twitter:url', content: META_URL }],
        ['meta', { property: 'twitter:title', content: META_TITLE }],
        ['meta', { property: 'twitter:description', content: META_DESCRIPTION }],
    ],

    themeConfig: {
        nav: [
            {
                text: 'Home',
                link: '/en/',
                activeMatch: '^/en/',
            },
            {
                text: 'Guide',
                link: '/en/guide/introduction',
                activeMatch: '^/en/guide/',
            },
            { text: 'API', link: '/en/api/overview', activeMatch: '^/en/api/' },
            {
                text: 'Links',
                items: [
                    {
                        text: 'Releaselog',
                        link: 'https://github.com/duke-git/lancet/releases',
                    },
                ],
            },
        ],

        sidebar: {
            '/en/': [
                {
                    text: 'Introduction',
                    items: [
                        {
                            text: 'What is Lancet？',
                            link: '/en/guide/introduction',
                        },
                        {
                            text: 'getting started',
                            link: '/en/guide/getting_started',
                        },
                    ],
                },
            ],
            '/en/api/': [
                {
                    text: 'overview',
                    items: [{ text: 'overview of API', link: '/en/api/overview' }],
                },
                {
                    text: 'packages',
                    items: [
                        { text: 'algorithm', link: '/en/api/packages/algorithm' },
                        { text: 'compare', link: '/en/api/packages/compare' },
                        { text: 'concurrency', link: '/en/api/packages/concurrency' },
                        { text: 'condition', link: '/en/api/packages/condition' },
                    ],
                },
            ],
        },
    },
}
