import React from 'react';
import Login from './views/login';
import Profile from './views/Profile';
function App() {
  return (
    <div>
      {localStorage.getItem('accessToken') ? (
        <Profile />
      ) : (
        <Login />
      )}
    </div>
  );
}

export default App;
