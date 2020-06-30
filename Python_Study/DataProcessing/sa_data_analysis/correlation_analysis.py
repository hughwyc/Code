import csv
import seaborn as sns
import matplotlib.pyplot as plt
import pandas as pd
import numpy as np

powder = "selfmade"  # 可选：selfmade  或  import
x_name = "speed"  # 可选 power 或 speed

# 读取csv至字典
csvFile = open('sa_' + powder + '_full_DF.csv', "r")
reader = csv.reader(csvFile)
power = []
speed = []
sa = []
for item in reader:
    # print(item[0],type(item[0]))
    if len(item[0]) > 0:
        power.append(int(float(item[1])))
        speed.append(float(item[2]))
        sa.append(float(item[3]))
csvFile.close()

# print(power)
# print(speed)
# print(sa)
data_full = [power, speed, sa]
data_full_arr = np.array(data_full)
data_full_arr = data_full_arr.T
data_full = data_full_arr.tolist()
data_name = ['power', 'speed', 'sa_value']
data = pd.DataFrame(columns=data_name, data=data_full)

# print(data)
rdata = data.corr()  # 查看数据间的相关系数
print(rdata)

# 作图

# # visualize the relationship between the features and the response using scatterplots
# sns.pairplot(data, x_vars=['power', 'speed'], y_vars='sa_value', kind="reg", size=5, aspect=0.7)
# # 通过加入一个参数kind='reg'，seaborn可以添加一条最佳拟合直线和95%的置信带。
sns.regplot(x=data[x_name], y=data["sa_value"])

if powder == "selfmade":
    poweder_chinese = "自制粉"
elif powder == "import":
    poweder_chinese = "进口粉"
plt.rcParams['font.sans-serif'] = ['SimHei']  # 配置显示中文，否则乱码
if x_name == "power":
    plt.xlim(135, 385)  # x轴刻度范围
    plt.xlabel('激光功率(w)')  # 添加x轴的名称
    plt.xticks(np.linspace(140, 380, 13))  # x轴的刻度值
    plt.title(poweder_chinese + '-激光功率与表面粗糙度散点图(回归线), R='+str(round(rdata.at[x_name, "sa_value"], 4)))  # 添加标题
elif x_name == "speed":
    plt.xlim(0.35, 1.25)  # x轴刻度范围
    plt.xlabel('扫描速度(m/s)')  # 添加x轴的名称
    plt.xticks(np.linspace(0.4, 1.2, 5))  # x轴的刻度值
    plt.title(poweder_chinese + '-扫描速度与表面粗糙度散点图(回归线), R='+str(round(rdata.at[x_name, "sa_value"], 4)))  # 添加标题
plt.ylim(0, 80)  # y轴刻度范围
plt.yticks(np.linspace(0, 80, 9))  # y轴的刻度值
plt.ylabel('表面粗糙度sa(μm)')  # 添加y轴的名称
# plt.legend(['a','b'])   #添加图例

plt.show()  # 注意必须加上这一句，否则无法显示。
