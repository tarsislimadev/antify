FROM tmvdl/dotnet:8-jammy

WORKDIR /app

COPY . .

RUN dotnet new tool-manifest --force # https://www.nuget.org/packages/dotnet-deb

RUN dotnet tool install --local dotnet-deb --version 0.1.226 # https://github.com/quamotion/dotnet-packaging

RUN dotnet deb install

RUN dotnet tool restore

RUN dotnet deb

CMD dotnet run
