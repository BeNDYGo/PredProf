document.addEventListener('DOMContentLoaded', async function () {
    const username = localStorage.getItem('username');

    if (!username) {
        document.getElementById('guest-message').style.display = 'block';
        return;
    }

    try {
        const response = await fetch(server + '/api/userInfo?username=' + username, {
            method: 'GET',
            headers: { 'Content-Type': 'application/json' }
        });

        if (!response.ok) {
            document.getElementById('guest-message').style.display = 'block';
            return;
        }

        const data = await response.json();
        const wins = data.wins || 0;
        const losses = data.losses || 0;
        const rating = data.rating || 0;
        const total = wins + losses;
        const winPercent = total > 0 ? Math.round((wins / total) * 100) : 0;

        document.getElementById('greeting').textContent = data.username;
        document.getElementById('wins-count').textContent = wins;
        document.getElementById('losses-count').textContent = losses;
        document.getElementById('win-percent').textContent = winPercent + '%';
        document.getElementById('rating-value').textContent = rating;
        document.getElementById('stats-section').style.display = 'block';

        if (total === 0) {
            document.getElementById('no-data').style.display = 'block';
            document.querySelector('.chart-container').style.display = 'none';
        } else {
            drawPieChart(wins, losses);
        }
    } catch (err) {
        document.getElementById('guest-message').style.display = 'block';
    }
});

function drawPieChart(wins, losses) {
    const total = wins + losses;
    const winRatio = wins / total;
    const size = 180;
    const cx = size / 2;
    const cy = size / 2;
    const r = 70;

    const winAngle = winRatio * 360;

    function polarToCartesian(angle) {
        const rad = (angle - 90) * Math.PI / 180;
        return {
            x: cx + r * Math.cos(rad),
            y: cy + r * Math.sin(rad)
        };
    }

    function arcPath(startAngle, endAngle) {
        const start = polarToCartesian(startAngle);
        const end = polarToCartesian(endAngle);
        const largeArc = (endAngle - startAngle) > 180 ? 1 : 0;
        return [
            'M', cx, cy,
            'L', start.x, start.y,
            'A', r, r, 0, largeArc, 1, end.x, end.y,
            'Z'
        ].join(' ');
    }

    let svg = `<svg width="${size}" height="${size}" viewBox="0 0 ${size} ${size}">`;

    if (wins === total) {
        svg += `<circle cx="${cx}" cy="${cy}" r="${r}" fill="#6c9"/>`;
    } else if (losses === total) {
        svg += `<circle cx="${cx}" cy="${cy}" r="${r}" fill="#e88"/>`;
    } else {
        svg += `<path d="${arcPath(0, winAngle)}" fill="#6c9"/>`;
        svg += `<path d="${arcPath(winAngle, 360)}" fill="#e88"/>`;
    }

    svg += '</svg>';
    document.getElementById('chart').innerHTML = svg;
}
