const server = 'https://deck-bedroom-peace-maximum.trycloudflare.com';

const username = localStorage.getItem('username');
const userLableInfo = document.getElementById("userInfo");

async function getUserInfo(username) {
    url = server + "/api/userInfo?username=" + username
    const response = await fetch(url, {
            method: "GET",
            headers: {
                'Content-Type': 'application/json',
            }
        });
    
    if (response.ok) {
        const data = await response.json();
        return data;
    } else {
        return null
    }
}

async function eloRender(){
    const data = await getUserInfo(username);

    if (data) {
        userLableInfo.innerHTML += `
            <div>
                ${data.username}: ${data.rating} ELO
            </div>
        `;
    }
}
eloRender();