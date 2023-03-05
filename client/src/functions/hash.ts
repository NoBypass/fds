import bcrypt from 'bcryptjs'

const saltRounds = 10

export const hash = (value: string): Promise<{ success: boolean, hash?: string, error?: string }> => {
    return new Promise((resolve, reject) => {
        bcrypt.genSalt(saltRounds, (err, salt) => {
            if (err) {
                reject({
                    success: false,
                    error: 'Error while generating salt: ' + err
                })
            } else {
                bcrypt.hash(value, salt, (err, hash) => {
                    if (err) {
                        reject({
                            success: false,
                            error: 'Error while generating hash: ' + err
                        })
                    } else {
                        resolve({
                            success: true,
                            hash
                        })
                    }
                })
            }
        })
    })
}

export const compareToHash = (value: string, hash: string): Promise<{ success: boolean, error?: string }> => {
    return new Promise((resolve, reject) => {
        bcrypt.compare(value, hash, (err, result) => {
            if (err) {
                reject({
                    success: false,
                    error: 'Error while comparing value with hash: ' + err
                })
            } else {
                resolve({
                    success: result
                })
            }
        })
    })
}