# Generated by Django 4.1 on 2023-02-21 20:45

from django.db import migrations, models
import django.db.models.deletion


class Migration(migrations.Migration):

    initial = True

    dependencies = []

    operations = [
        migrations.CreateModel(
            name="Student",
            fields=[
                (
                    "id",
                    models.BigAutoField(
                        auto_created=True,
                        primary_key=True,
                        serialize=False,
                        verbose_name="ID",
                    ),
                ),
                (
                    "first_name",
                    models.CharField(max_length=60, verbose_name="First Name"),
                ),
                (
                    "last_name",
                    models.CharField(max_length=60, verbose_name="Last Name"),
                ),
                (
                    "reg_no",
                    models.CharField(
                        default="REG", max_length=60, verbose_name="Registration Number"
                    ),
                ),
            ],
        ),
        migrations.CreateModel(
            name="Tutor",
            fields=[
                (
                    "id",
                    models.BigAutoField(
                        auto_created=True,
                        primary_key=True,
                        serialize=False,
                        verbose_name="ID",
                    ),
                ),
                ("first_name", models.CharField(max_length=60)),
                ("last_name", models.CharField(max_length=60)),
            ],
        ),
        migrations.CreateModel(
            name="Subject",
            fields=[
                (
                    "subject",
                    models.CharField(
                        default="Subject",
                        max_length=60,
                        primary_key=True,
                        serialize=False,
                    ),
                ),
                ("student", models.ManyToManyField(blank=True, to="workspace.student")),
                ("tutor", models.ManyToManyField(to="workspace.tutor")),
            ],
        ),
        migrations.AddField(
            model_name="student",
            name="tutor",
            field=models.ManyToManyField(blank=True, to="workspace.tutor"),
        ),
        migrations.CreateModel(
            name="Assignment",
            fields=[
                (
                    "id",
                    models.BigAutoField(
                        auto_created=True,
                        primary_key=True,
                        serialize=False,
                        verbose_name="ID",
                    ),
                ),
                ("topic", models.CharField(max_length=60)),
                ("discription", models.TextField(blank=True)),
                ("due_date", models.DateTimeField(verbose_name="Due Date")),
                ("student", models.ManyToManyField(blank=True, to="workspace.student")),
                (
                    "subject",
                    models.ForeignKey(
                        default="Subject",
                        on_delete=django.db.models.deletion.CASCADE,
                        to="workspace.subject",
                    ),
                ),
            ],
        ),
    ]
