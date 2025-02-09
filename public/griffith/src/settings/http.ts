export const URL = "http://localhost:8080"

export const callAPI = async (path: string, method: string, body: any, noAuth: boolean = true): Promise<any> => {
    let headers = undefined
    if (!noAuth) {
        const token = localStorage.getItem('accessToken')
        if (!token) {
            return Error('No token found')
        }
        headers = {
            'Authorization': `${token}`
        }
    }
    let bodytopass = null

    if (body) {
        bodytopass = JSON.stringify(body)
    } else{
        bodytopass = undefined
    }

    const response = await fetch(`${URL}${path}`, {
        method: method,
        body: bodytopass,
        headers: headers
    })

    if (response.status === 200) {
      const data = await response.json()
      return data
    } else if (response.status === 401) {
        const resp = await refreshToken()
        if (resp instanceof Error) {
            return resp
        }
        return callAPI(path, method, body, noAuth)
    } else {
        const data = await response.json()
        return Error(data)
    }
}

const refreshToken = async () => {
    console.log("Refreshing token")
    const headers = {
        'Authorization': `${localStorage.getItem('refreshToken')}`
    }
    const body = {
        'refreshToken': localStorage.getItem('refreshToken')
    }
    console.log(body)
    const response = await fetch(`${URL}/auth/refresh`, {
        method: 'POST', 
        headers: headers,
        body: JSON.stringify(body)
    })
    if (response.status === 200) {
        const data = await response.json()
        localStorage.setItem('accessToken', data.accessToken) 
    } else {
        const data = await response.json()
        console.log(data)
        localStorage.removeItem('accessToken')
        localStorage.removeItem('refreshToken')
        window.location.href = '/'
        return Error('Failed to refresh token')
    }
}