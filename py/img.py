from PIL import Image, ImageDraw

if __name__ == '__main__':
    # 创建一个白色背景的300 x 300图像
    img = Image.new('RGB', (300, 300), color='white')

    # 绘制卡通人物
    draw = ImageDraw.Draw(img)
    draw.ellipse((75, 30, 225, 150), fill='yellow')  # 脸
    draw.ellipse((90, 50, 120, 80), fill='black')  # 左眼
    draw.ellipse((180, 50, 210, 80), fill='black')  # 右眼
    draw.polygon([(150, 100), (120, 130), (180, 130)], fill='red')  # 嘴巴

    # 显示图像
    img.show()
