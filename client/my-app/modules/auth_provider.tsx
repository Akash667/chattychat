import { createContext, useState } from 'react'

export type UserInfo = {
    username: string
    id: string
}

export const AuthContext = createContext<
    {
        authenticated: boolean, setAuthenticated: (auth: boolean) => void, user: UserInfo, setUser: (user: UserInfo) => void
    }>(
        {
            authenticated: false,
            setAuthenticated: () => { },
            user: { username: '', id: '' },
            setUser: () => { }
        }
    )

const AuthContextProvider = ({ children: React.ReactNode }) => {

    const [authenticated, setAuthenticated] = useState(false)
    const [user, setUser] = useState<UserInfo>({ username: '', id: '' })
    return (
        <div>AuthContextProvider</div>
    )
}

export default AuthContextProvider
