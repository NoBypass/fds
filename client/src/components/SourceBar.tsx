import * as React from 'react';
import {useEffect, useState} from "react";
import fastFetch from "../functions/fast-fetch";

export default function SourceBar() {
    const defaultData = {name: 'Loading...', url: '#'}
    const [data, setData] = useState<any>(defaultData)          //TODO create type for data

    const name = 'NoBypass'
    const repoName = 'fds'
    const options = {
        headers: {
            Authorization: `Bearer ${"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"}`,
            'User-Agent': 'My-App',
        },
    }

    useEffect(() => {
        if (data === defaultData) {
            const callFn = async () => {
                console.log(import.meta.env.VITE_GH_TOKEN)

                const [repo, commits, contributors] = await Promise.all([
                    fastFetch(`https://api.github.com/repos/${name}/${repoName}`, options),
                    fastFetch(`https://api.github.com/repos/${name}/${repoName}/commits`, options),
                    fastFetch(`https://api.github.com/repos/${name}/${repoName}/contributors`, options),
                ])
                const tags = await fastFetch(`https://api.github.com/repos/${name}/${repoName}/commits/${commits[0].sha}/refs/tags`, options)           // TODO: make tags to be iterated through

                console.log([repo, commits, contributors, tags])
                setData({
                    name: repo.name,
                    url: repo.url,
                    version: `v-${tags}p-${0}d-${commits[0].committer.date}` // TODO: remove 'tags' use latest tag instead | instead of 0 user commits since last tag
                })
            }
            callFn()
        }
    }, [data])

    return (
        <section>
            <a href={data?.url}>{data?.name}</a>
            <a></a>
        </section>
    )
}