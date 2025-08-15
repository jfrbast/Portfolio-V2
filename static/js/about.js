class RouladeSprite {
    constructor(containerId) {
        this.container = document.getElementById(containerId);
        this.currentFrame = 1;
        this.totalFrames = 8;
        this.isAnimating = false;
        this.animationSpeed = 100;

        this.createSpriteElement();
    }

    createSpriteElement() {
        this.sprite = document.createElement('img');
        this.sprite.src = '/static/sprites/roulade/roll1.png';
        this.sprite.alt = 'Roulade animation';

        this.container.appendChild(this.sprite);
    }

    animate() {
        if (this.isAnimating) return;

        this.isAnimating = true;
        this.animationInterval = setInterval(() => {
            this.currentFrame++;
            if (this.currentFrame > this.totalFrames) {
                this.currentFrame = 1;
            }
            this.sprite.src = `/static/sprites/roulade/roll${this.currentFrame}.png`;
        }, this.animationSpeed);
    }

    stop() {
        if (this.animationInterval) {
            clearInterval(this.animationInterval);
            this.isAnimating = false;
        }
    }
}

class CowSprite {
    constructor(containerId) {
        this.container = document.getElementById(containerId);
        this.currentFrame = 1;
        this.totalFrames = 8;
        this.isAnimating = false;
        this.animationSpeed = 150;

        this.createSpriteElement();
    }

    createSpriteElement() {
        this.sprite = document.createElement('img');
        this.sprite.src = '/static/sprites/cow/cow1.png';
        this.sprite.alt = 'Cow animation';

        this.container.appendChild(this.sprite);
    }

    animate() {
        if (this.isAnimating) return;

        this.isAnimating = true;
        this.animationInterval = setInterval(() => {
            this.currentFrame++;
            if (this.currentFrame > this.totalFrames) {
                this.currentFrame = 1;
            }
            this.sprite.src = `/static/sprites/cow/cow${this.currentFrame}.png`;
        }, this.animationSpeed);
    }

    stop() {
        if (this.animationInterval) {
            clearInterval(this.animationInterval);
            this.isAnimating = false;
        }
    }
}

document.addEventListener('DOMContentLoaded', function() {
    const roulade = new RouladeSprite('sprite-container');
    const cow = new CowSprite('cow-container');

    roulade.animate();
    cow.animate();
});