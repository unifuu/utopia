const canvas = document.getElementById('gameCanvas');
const ctx = canvas.getContext('2d');

const player = {
    x: 50,
    y: 50,
    width: 50,
    height: 50,
    speed: 10,
    color: 'purple'
};

function drawPlayer() {
    ctx.fillStyle = player.color;
    ctx.fillRect(player.x, player.y, player.width, player.height);
}

function clearCanvas() {
    ctx.clearRect(0, 0, canvas.width, canvas.height);
}

function updatePlayerPosition(e) {
    switch(e.key) {
        case 'ArrowUp':
            player.y -= player.speed;
            break;
        case 'ArrowDown':
            player.y += player.speed;
            break;
        case 'ArrowLeft':
            player.x -= player.speed;
            break;
        case 'ArrowRight':
            player.x += player.speed;
            break;
    }
    keepPlayerInBounds();
}

function keepPlayerInBounds() {
    if (player.x < 0) player.x = 0;
    if (player.y < 0) player.y = 0;
    if (player.x + player.width > canvas.width) player.x = canvas.width - player.width;
    if (player.y + player.height > canvas.height) player.y = canvas.height - player.height;
}

function gameLoop() {
    clearCanvas();
    drawPlayer();
    requestAnimationFrame(gameLoop);
}

document.addEventListener('keydown', updatePlayerPosition);

gameLoop();