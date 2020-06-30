import csv
import pandas as pd
from matplotlib import pyplot as plt
import numpy as np
from scipy import interpolate
import seaborn as sns
from matplotlib import cm

powder = "selfmade"  # 可选：selfmade  或  import

# 读取csv至字典
csvFile = open("sa_" + powder + ".csv", "r")
reader = csv.reader(csvFile)

result = {}  # 用于存储读取到的数据
for item in reader:
    if len(item) > 0:
        result[item[0]] = item[1]

csvFile.close()

# for k, v in result.items():
#     print("key: " + k + ", value: " + str(v))

ave_dict = {}  # 用于存储求平均数后的数据
for number in range(1, 66):
    sum = 0.0
    count = 0  # 累计同一组的次数
    name = ''  # 用于存储试样的编号
    for k, v in result.items():
        if len(k) != 0 and int(k[0:3]) == number:
            name = k[:-2]
            count += 1
            sum += float(v)
    if count != 0:
        ave_dict[name] = round(sum / count, 3)

power = []  # 用于保存不重复的激光功率
speed = []
power_full = []  # 用于保存所有激光功率，包括每5组重复的
speed_full = []
sa = []

for k, v in ave_dict.items():
    # print("key: " + k + ", value: " + str(v)) #  key: 065-380-1.2, value: 41.265
    # print("power=", k[4:7],"speed=", k[-3:], "sa=",str(v))
    power_full.append(int(k[4:7]))
    speed_full.append(float(k[-3:]))
    if int(k[4:7]) not in power:
        power.append(int(k[4:7]))
    if float(k[-3:]) not in speed:
        speed.append(float(k[-3:]))
    sa.append(v)
# print(power)
# print(speed)
# print(sa)

# 借助np.array, 将横向列表，转成纵向列表
data_full = [power_full, speed_full, sa]
data_full_arr = np.array(data_full)
data_full_arr = data_full_arr.T
data_full = data_full_arr.tolist()
# print(data_full)
data_name = ['power', 'speed', 'sa_value']
data_full_df = pd.DataFrame(columns=data_name, data=data_full)
print(data_full_df)
data_full_df.to_csv('sa_'+powder + '_full_DF.csv', encoding='gbk')  # 将“全数据”输出成csv格式

# 将sa一维向量转换成二维数组
row = []
sa_2d = []
for i in range(len(sa)):
    row.append(sa[i])
    if (i + 1) % 5 == 0:
        sa_2d.append(row)
        row = []
# print(sa_2d)

# 将三组数据转成DataFrame数据输出到csv中，分别对应行名、列名、sa数据
data = pd.DataFrame(columns=speed, index=power, data=sa_2d)  # 将数据转换为DataFrame类型
data.to_csv('sa_'+powder + '_DF.csv', encoding='gbk')  # 输出成csv格式
print(data)

row = data.index.values.tolist()  # 行名称
col = data.columns.values.tolist()  # 列名称
row.reverse()  # 直接在原来的列表里面将元素进行逆序排列
# col.reverse()

data = data.loc[row, col]
print(data)
# print(data.values)

# 画图
# fig = plt.figure() # figsize=(15, 15)
fig = plt.gcf()
fig.set_size_inches(12, 12)
ax = fig.add_subplot(111, projection='3d')
# ax.set_figure
# 调整观察角度和方位角。这里将俯仰角设为60度，把方位角调整为35度
ax.view_init(14, -50)
# 设置xyz坐标轴范围
plt.xlim(1.25, 0.35)
plt.ylim(385, 135)
ax.set_zlim((0, 80))
# 设置xy轴刻度显示
plt.xticks([0.4, 0.6, 0.8, 1.0, 1.2])
plt.yticks(np.linspace(380, 140, 13))
plt.rcParams['font.sans-serif'] = ['SimHei']  # 配置显示中文，否则乱码
plt.xlabel('扫描速度(m/s)')
plt.ylabel('激光功率(w)')
ax.set_zlabel("粗糙度sa(μm)")
if powder=="selfmade":
    poweder_chinese = "自制粉"
elif powder == "import":
    poweder_chinese = "进口粉"
plt.title(poweder_chinese+"打印样品粗糙度与激光功率、扫描速度的关系")
# 定义三维数据
x = np.array(row)
y = np.array(col)
z = np.array(data.values)
# print(x)
# print(y)
# print(z)
z = z.T
# print(z)
X, Y = np.meshgrid(x, y)

# 插值
f = interpolate.interp2d(x, y, z, kind='cubic')
xnew = np.linspace(140, 380, 100)
ynew = np.linspace(0.4, 1.2, 100)
znew = f(xnew, ynew)

# 修改x,y，z输入画图函数前的shape
xx1, yy1 = np.meshgrid(xnew, ynew)
newshape = (xx1.shape[0]) * (xx1.shape[0])
y_input = xx1.reshape(newshape)
x_input = yy1.reshape(newshape)
z_input = znew.reshape(newshape)

# 作图
# sns.set(style='white')
ax.plot_trisurf(x_input, y_input, z_input, cmap=cm.coolwarm)
ax.scatter3D(Y, X, z + 0.5, cmap='Blues')  # 绘制散点图

plt.show()
