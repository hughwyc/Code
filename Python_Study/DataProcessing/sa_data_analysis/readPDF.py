# -*- coding: utf-8 -*-
"""
Created on Sat Jun 13 10:53:47 2020

@author: Hugh
"""

from io import StringIO
from io import open
from pdfminer.converter import TextConverter
from pdfminer.layout import LAParams
from pdfminer.pdfinterp import PDFResourceManager, process_pdf
import os
import csv


def read_pdf(pdf):
    # resource manager
    rsrcmgr = PDFResourceManager()
    retstr = StringIO()
    laparams = LAParams()
    # device
    device = TextConverter(rsrcmgr, retstr, laparams=laparams)
    process_pdf(rsrcmgr, device, pdf)
    device.close()
    content = retstr.getvalue()
    retstr.close()
    # 获取所有行
    pdf_lines = str(content).split("\n")
    return pdf_lines


if __name__ == '__main__':

    sa_dict = {}  # 用于存储目标值sa

    # 获取路径下所有文件的全路径
    file_dir = "D:/Documents/科研数据/M1表面粗糙度/TC4-自制粉-2020-06-06打印/"
    file_path = []
    for file in os.listdir(file_dir):
        filename = os.path.join(file_dir, file)
        result = []
        # 读取每个子文件pdf内容
        with open(filename, "rb") as my_pdf:
            pdf_lines = read_pdf(my_pdf)
        for item in pdf_lines:
            # 删除每个元素中的空格，和空元素
            if len(item) > 0:
                result.append(item.replace(' ', ''))
        # print(result)
        loc = result.index('Visualization')  # 'Visualization'的下一个元素就是sa的值
        # 找到目标数据sa的值，存入字典sa_dict
        sa_dict[file[:-4]] = float(result[loc + 1][:-2])  # [:-2]，不读取um单位

    # for k, v in sa_dict.items():
    #     print("key: " + k + ", value: " + str(v))

    # 将读取到的sa字典输出到csv中
    with open('sa_selfmade.csv', 'w') as csv_file:
        writer = csv.writer(csv_file)
        for key, value in sa_dict.items():
            writer.writerow([key, value])
    # 注意此时的数据是三列，并未做取均值操作，且不是DataFrame类型