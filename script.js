const slides = document.querySelectorAll('.slide');
const intervalTime = 5000;
let slideInterval;

const nextSlide = () => {
    // 获取当前活动幻灯片
    const activeSlide = document.querySelector('.active');
    // 获取下一个幻灯片
    let nextSlide = activeSlide.nextElementSibling;
    // 如果下一个幻灯片不存在，则获取第一个幻灯片
    if (!nextSlide) {
        nextSlide = slides[0];
    }
    // 添加活动类
    activeSlide.classList.remove('active');
    nextSlide.classList.add('active');
}

const startSlide = () => {
    slideInterval = setInterval(() => {
        nextSlide();
    }, intervalTime);
}

const stopSlide = () => {
    clearInterval(slideInterval);
}

// 在侧边栏导航中添加新卡片
const addCard = (title, content) => {
    const sidebar = document.querySelector('.sidebar');
    const card = document.createElement('div');
    card.classList.add('sidebar-card');
    const cardTitle = document.createElement('h3');
    cardTitle.innerText = title;
    const cardContent = document.createElement('p');
    cardContent.innerText = content;
    const cardLinks = document.createElement('ul');
    const link1 = document.createElement('li');
    const link2 = document.createElement('li');
    const link3 = document.createElement('li');
    link1.innerHTML = '<a href="#">链接1</a>';
    link2.innerHTML = '<a href="#">链接2</a>';
    link3.innerHTML = '<a href="#">链接3</a>';
    cardLinks.appendChild(link1);
    cardLinks.appendChild(link2);
    cardLinks.appendChild(link3);
    card.appendChild(cardTitle);
    card.appendChild(cardContent);
    card.appendChild(cardLinks);
    sidebar.appendChild(card);
}
