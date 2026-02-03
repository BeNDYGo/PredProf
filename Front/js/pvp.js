const server = 'http://localhost:8080';
const username = localStorage.getItem('username');
const userLableInfo = document.getElementById("userInfo");

async function getUserInfo(username) {
    url = server + "/userInfo?username=" + username
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
async function renderUser() {
    const data = await getUserInfo(username);

    if (!data) return;

    userLableInfo.innerHTML += `
        <div>
            ${data.username} — рейтинг: ${data.rating}
        </div>
    `;
}

renderUser();