const query = async <T>(req: {
    query: string,
    variables: any,
    operationName?: string,
    token?: string,
}): Promise<{ status: number, data: T }> => {
    const res = await fetch('http://localhost:8080/graphql', { // TODO: use env variable
        method: 'POST',
        headers: {
            'Content-Type': 'application/json',
        },
        body: JSON.stringify({
            operationName: req.operationName || '',
            query: req.query.replace(/\s+/g, ''),
            variables: req.variables || {},
        }),
    })

    return new Promise((resolve, reject) => {
        res.json()
            .then((data) => {
                resolve({
                    status: res.status,
                    data: (data as {
                        data: {
                            [key: string]: T
                        }}).data[req.query.split('{')[1].split('(')[0].trim()],
                })
            })
            .catch(async (e) => {
                reject({
                    status: res.status,
                    msg: e,
                    error: e,
                })
            })
    })
}

export default query
